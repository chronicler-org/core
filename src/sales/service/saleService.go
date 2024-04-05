package saleService

import (
	"errors"
	"time"

	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	productService "github.com/chronicler-org/core/src/product/service"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleEnum "github.com/chronicler-org/core/src/sales/enum"
	saleExceptionMessage "github.com/chronicler-org/core/src/sales/messages"
	salesModel "github.com/chronicler-org/core/src/sales/model"
	salesRepository "github.com/chronicler-org/core/src/sales/repository"
)

type SaleService struct {
	saleRepository      *salesRepository.SaleRepository
	saleItemRepository  *salesRepository.SaleItemRepository
	customerCareService *customerCareService.CustomerCareService
	productService      *productService.ProductService
}

func InitSaleService(
	saleRepository *salesRepository.SaleRepository,
	saleItemRepository *salesRepository.SaleItemRepository,
	productService *productService.ProductService,
	customerCareService *customerCareService.CustomerCareService,
) *SaleService {
	return &SaleService{
		saleRepository:      saleRepository,
		saleItemRepository:  saleItemRepository,
		customerCareService: customerCareService,
		productService:      productService,
	}
}

func (service *SaleService) FindSaleByID(id string) (salesModel.Sale, error) {
	result, err := service.saleRepository.FindOneByField("CustomerCareID", id, "CustomerCare")
	sale, _ := result.(*salesModel.Sale)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *sale, appException.NotFoundException(saleExceptionMessage.SALE_NOT_FOUND)
	}

	return *sale, nil
}

func (service *SaleService) CreateSale(
	dto salesDTO.CreateSaleDTO,
) (salesModel.Sale, error) {
	customerCareExists, err := service.customerCareService.FindCustomerCareByID(dto.CustomerCareID)

	if err != nil {
		return salesModel.Sale{}, err
	}

	transaction := service.saleRepository.BeginTransaction()

	// Create the sale record
	saleModel := salesModel.Sale{
		CustomerCareID: customerCareExists.ID,
		Status:         dto.Status,
		PaymentMethod:  dto.PaymentMethod,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	err = service.saleRepository.CreateWithTransaction(transaction, saleModel)
	if err != nil {
		transaction.Rollback()
		return salesModel.Sale{}, err
	}

	// Calculate the total value of the sale and create sale items
	var totalValue float32
	for _, itemDTO := range dto.SalesItems {
		saleItem := salesModel.SaleItem{
			SaleID:    saleModel.CustomerCareID,
			Quantity:  itemDTO.Quantity,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		productHasStockAvailable, err := service.productService.ValidateStock(itemDTO.ProductID, itemDTO.Quantity)
		if err != nil {
			transaction.Rollback()
			return salesModel.Sale{}, err
		}
		totalValue += productHasStockAvailable.Value * float32(itemDTO.Quantity)
		saleItem.ProductID = productHasStockAvailable.ID

		err = service.saleItemRepository.CreateWithTransaction(transaction, saleItem)
		if err != nil {
			transaction.Rollback()
			return salesModel.Sale{}, err
		}

		_, err = service.productService.DebitStock(itemDTO.ProductID, itemDTO.Quantity, transaction)
		if err != nil {
			transaction.Rollback()
			return salesModel.Sale{}, err
		}
	}

	// Update the sale record with the total value
	saleModel.TotalValue = totalValue
	saleModel.UpdatedAt = time.Now()
	err = service.saleRepository.UpdateWithTransaction(transaction, saleModel)
	if err != nil {
		transaction.Rollback()
		return salesModel.Sale{}, err
	}

	transaction.Commit()
	saleModel.CustomerCare = customerCareExists
	return saleModel, nil
}

func (service *SaleService) FindAllSales(dto salesDTO.QuerySalesDTO) (int64, []salesModel.Sale, error) {
	var sales []salesModel.Sale

	count, err := service.saleRepository.FindAll(dto, &sales, "CustomerCare")
	if err != nil {
		return 0, nil, err
	}

	return count, sales, err
}

func (service *SaleService) UpdateSale(updateSaleDTO salesDTO.UpdateSaleDTO, id string) (salesModel.Sale, error) {
	sale, err := service.FindSaleByID(id)
	if err != nil {
		return salesModel.Sale{}, err
	}

	switch sale.Status {
	case string(saleEnum.AGUARDANDO_PAGAMENTO):
		switch updateSaleDTO.Transition {

		case string(saleEnum.PAGAMENTO_CONFIRMADO):
			sale.Status = string(saleEnum.COMPRA_CONFIRMADA)

		case string(saleEnum.CANCELAR_COMPRA):
			sale.Status = string(saleEnum.COMPRA_CANCELADA)

		default:
			return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
		}

	case string(saleEnum.COMPRA_CONFIRMADA):
		switch updateSaleDTO.Transition {

		case string(saleEnum.CONCLUIR_COMPRA):
			sale.Status = string(saleEnum.COMPRA_CONCLUIDA)

		case string(saleEnum.CANCELAR_COMPRA):
			sale.Status = string(saleEnum.COMPRA_CANCELADA)

		default:
			return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
		}

	default:
		return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
	}

	transaction := service.saleRepository.BeginTransaction()

	if sale.Status == string(saleEnum.COMPRA_CANCELADA) {
		var items []salesModel.SaleItem
		_, err := service.saleItemRepository.FindAll(
			salesDTO.QuerySaleItemsDTO{
				SaleID: sale.CustomerCareID.String(),
			},
			&items,
			"CustomerCare",
		)
		if err != nil {
			return salesModel.Sale{}, err
		}

		for _, item := range items {
			_, err := service.productService.CreditStock(item.ProductID.String(), item.Quantity, transaction)
			if err != nil {
				transaction.Rollback()
				return salesModel.Sale{}, err
			}
		}

		service.saleItemRepository.DeleteWithTransaction(transaction, "SaleID", sale.CustomerCareID.String())
	}

	sale.UpdatedAt = time.Now()

	err = service.saleRepository.UpdateWithTransaction(transaction, sale)
	if err != nil {
		transaction.Rollback()
		return salesModel.Sale{}, err
	}

	transaction.Commit()
	return sale, err
}

func (service *SaleService) DeleteSale(id string) (salesModel.Sale, error) {
	sale, err := service.FindSaleByID(id)
	if err != nil {
		return salesModel.Sale{}, err
	}

	err = service.saleRepository.Delete("CustomerCareID", id)
	if err != nil {
		return salesModel.Sale{}, err
	}

	return sale, err
}
