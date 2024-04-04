package saleService

import (
	"errors"
	"time"

	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleExceptionMessage "github.com/chronicler-org/core/src/sales/messages"
	salesModel "github.com/chronicler-org/core/src/sales/model"
	salesRepository "github.com/chronicler-org/core/src/sales/repository"
)

type SaleService struct {
	saleRepository      *salesRepository.SaleRepository
	saleItemRepository  *salesRepository.SaleItemRepository
	customerCareService *customerCareService.CustomerCareService
}

func InitSaleService(
	saleRepository *salesRepository.SaleRepository,
	saleItemRepository *salesRepository.SaleItemRepository,
	customerCareService *customerCareService.CustomerCareService,
) *SaleService {
	return &SaleService{
		saleRepository:      saleRepository,
		customerCareService: customerCareService,
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

	model := salesModel.Sale{
		CustomerCareID: customerCareExists.ID,
		TotalValue:     dto.TotalValue,
		Status:         dto.Status,
		PaymentMethod:  dto.PaymentMethod,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = service.saleRepository.Create(model)
	if err != nil {
		return salesModel.Sale{}, err
	}

	model.CustomerCare = customerCareExists
	return model, err
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
