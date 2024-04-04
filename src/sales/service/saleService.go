package saleService

import (
	"errors"
	"time"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleExceptionMessage "github.com/chronicler-org/core/src/sales/messages"
	salesModel "github.com/chronicler-org/core/src/sales/model"
	salesRepository "github.com/chronicler-org/core/src/sales/repository"
	"gorm.io/gorm"
)

type SaleService struct {
	salesRepository     *salesRepository.SalesRepository
	customerCareService *customerCareService.CustomerCareService
}

func InitSaleService(
	salesRepository *salesRepository.SalesRepository,
	customerCareService *customerCareService.CustomerCareService,
) *SaleService {
	return &SaleService{
		salesRepository:     salesRepository,
		customerCareService: customerCareService,
	}
}

func (service *SaleService) FindSaleByID(id string) (salesModel.Sales, error) {
	result, err := service.salesRepository.FindOneByField("CustomerCareID", id, "CustomerCare")
	sale, _ := result.(*salesModel.Sales)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *sale, appException.NotFoundException(saleExceptionMessage.SALE_NOT_FOUND)
	}

	return *sale, nil
}

func (service *SaleService) CreateSale(
	dto salesDTO.CreateSaleDTO,
	customerCareID string,
) (salesModel.Sales, error) {
	customerCareExists, err := service.customerCareService.FindCustomerCareByID(customerCareID)

	if err != nil {
		return salesModel.Sales{}, err
	}

	model := salesModel.Sales{
		CustomerCareID: customerCareExists.ID,
		TotalValue:     dto.TotalValue,
		Status:         dto.Status,
		PaymentMethod:  dto.PaymentMethod,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = service.salesRepository.Create(model)
	if err != nil {
		return salesModel.Sales{}, err
	}

	model.CustomerCare = customerCareExists
	return model, err
}

func (service *SaleService) FindAllSales(dto salesDTO.QuerySalesDTO) (int64, []salesModel.Sales, error) {
	var sales []salesModel.Sales

	count, err := service.salesRepository.FindAll(dto, &sales, "CustomerCare")
	if err != nil {
		return 0, nil, err
	}

	return count, sales, err
}

func (service *SaleService) UpdateSale(dto salesDTO.UpdateSaleDTO, id string) (salesModel.Sales, error) {
	sale, err := service.FindSaleByID(id)
	if err != nil {
		return salesModel.Sales{}, err
	}

	appUtil.UpdateModelFromDTO(&sale, &dto)

	sale.UpdatedAt = time.Now()

	err = service.salesRepository.Update(sale)
	if err != nil {
		return salesModel.Sales{}, nil
	}

	return sale, err
}

func (service *SaleService) DeleteSale(id string) (salesModel.Sales, error) {
	sale, err := service.FindSaleByID(id)
	if err != nil {
		return salesModel.Sales{}, err
	}

	err = service.salesRepository.Delete("CustomerCareID", id)
	if err != nil {
		return salesModel.Sales{}, err
	}

	return sale, err
}
