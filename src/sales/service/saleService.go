package saleService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	productEnum "github.com/chronicler-org/core/src/product/enum"
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
		Status:         saleEnum.AWAITING_PAYMENT,
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
	case saleEnum.AWAITING_PAYMENT:
		switch updateSaleDTO.Transition {

		case string(saleEnum.PAGAMENTO_CONFIRMADO):
			sale.Status = saleEnum.PURCHASE_CONFIRMED

		case string(saleEnum.CANCELAR_COMPRA):
			sale.Status = saleEnum.CANCELLED_PURCHASE

		default:
			return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
		}

	case saleEnum.PURCHASE_CONFIRMED:
		switch updateSaleDTO.Transition {

		case string(saleEnum.CONCLUIR_COMPRA):
			sale.Status = saleEnum.PURCHASE_COMPLETED

		case string(saleEnum.CANCELAR_COMPRA):
			sale.Status = saleEnum.CANCELLED_PURCHASE

		default:
			return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
		}

	default:
		return sale, appException.ConflictException(saleExceptionMessage.INVALID_TRANSITION)
	}

	transaction := service.saleRepository.BeginTransaction()

	if sale.Status == saleEnum.CANCELLED_PURCHASE {
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

func (service *SaleService) GetSaleProductsSummary(dto salesDTO.QuerySalesProductSummaryDTO) (interface{}, int64, error) {
	produtsSummary := []struct {
		ProductID     uuid.UUID                 `json:"product_id"`
		Model         productEnum.ClothingModel `json:"model"`
		TotalQuantity int64                     `json:"total_quantity"`
	}{}

	totalCount, err := service.saleItemRepository.GetSaleProductSummary(dto, &produtsSummary)
	if err != nil {
		return nil, 0, err
	}

	return produtsSummary, totalCount, nil
}

func (service *SaleService) GetProductQuantitySoldVariation() (any, error) {
	productQuantitySoldVariation := struct {
		PercentVariation  float64 `json:"percent_variation"`
		TotalCurrentMonth int64   `json:"total_current_month"`
	}{}

	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()
	currentMonthTotalQuantity, err :=
		service.saleItemRepository.GetTotalQuantitySoldByCreatedMonth(currentMonth, currentYear)
	if err != nil {
		return productQuantitySoldVariation, err
	}

	lastMonth, lastYear := appUtil.GetLastMonth()
	lastMonthTotalQuantity, err :=
		service.saleItemRepository.GetTotalQuantitySoldByCreatedMonth(lastMonth, lastYear)
	if err != nil {
		return productQuantitySoldVariation, err
	}

	productQuantitySoldVariation.TotalCurrentMonth = currentMonthTotalQuantity
	productQuantitySoldVariation.PercentVariation =
		appUtil.CalculatePercentVariation(float64(currentMonthTotalQuantity), float64(lastMonthTotalQuantity))

	return productQuantitySoldVariation, nil
}

func (service *SaleService) FindAllSaleItems(dto salesDTO.QuerySaleItemDTO) ([]salesModel.SaleItem, int64, error) {
	var saleItems []salesModel.SaleItem

	totalCount, err := service.saleItemRepository.FindAll(dto, &saleItems,
		"Sale", "Product", "Sale.CustomerCare",
		"Sale.CustomerCare.Customer", "Sale.CustomerCare.Team",
		"Sale.CustomerCare.Customer.Address",
	)
	if err != nil {
		return nil, 0, err
	}

	return saleItems, totalCount, nil
}
