package saleService

import (
	"errors"
	"time"

	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	productService "github.com/chronicler-org/core/src/product/service"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
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
	customerCareID string,
) (salesModel.Sale, error) {
	customerCareExists, err := service.customerCareService.FindCustomerCareByID(customerCareID)

	if err != nil {
		return salesModel.Sale{}, err
	}

	saleModel := salesModel.Sale{
		CustomerCareID: customerCareExists.ID,
		Status:         dto.Status,
		PaymentMethod:  dto.PaymentMethod,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	var totalValue float32
	salesItems := []salesModel.SaleItem{}
	for _, itemDTO := range dto.SalesItems {
		saleItem := salesModel.SaleItem{
			SaleID:    saleModel.CustomerCareID,
			Quantity:  itemDTO.Quantity,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		product, err := service.productService.FindProductByID(itemDTO.ProductID)
		totalValue += product.Value * float32(itemDTO.Quantity)
		if err != nil {
			return salesModel.Sale{}, err
		}
		saleItem.ProductID = product.ID
		salesItems = append(salesItems, saleItem)
	}

	saleModel.TotalValue = totalValue
	err = service.saleRepository.Create(saleModel)
	if err != nil {
		return salesModel.Sale{}, err
	}

	for _, saleItem := range salesItems {
		err := service.saleItemRepository.Create(saleItem)
		if err != nil {
			service.saleRepository.Delete("CustomerCareID", saleModel.CustomerCareID.String())
			return salesModel.Sale{}, err
		}
	}

	saleModel.CustomerCare = customerCareExists
	return saleModel, err
}

func (service *SaleService) FindAllSales(dto salesDTO.QuerySalesDTO) (int64, []salesModel.Sale, error) {
	var sales []salesModel.Sale

	count, err := service.saleRepository.FindAll(dto, &sales, "CustomerCare")
	if err != nil {
		return 0, nil, err
	}

	return count, sales, err
}

func (service *SaleService) UpdateSale(dto salesDTO.UpdateSaleDTO, id string) (salesModel.Sale, error) {
	sale, err := service.FindSaleByID(id)
	if err != nil {
		return salesModel.Sale{}, err
	}

	appUtil.UpdateModelFromDTO(&sale, &dto)

	sale.UpdatedAt = time.Now()

	err = service.saleRepository.Update(sale)
	if err != nil {
		return salesModel.Sale{}, nil
	}

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
