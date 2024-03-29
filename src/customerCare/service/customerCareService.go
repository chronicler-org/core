package customerCareService

import (
	"errors"
	"time"

	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	customerService "github.com/chronicler-org/core/src/customer/service"
	customerCareDTO "github.com/chronicler-org/core/src/customerCare/dto"
	customerCareExceptionMessage "github.com/chronicler-org/core/src/customerCare/messages"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	customerCareRepository "github.com/chronicler-org/core/src/customerCare/repository"
	teamService "github.com/chronicler-org/core/src/team/service"
)

type CustomerCareService struct {
	customerCareRepository *customerCareRepository.CustomerCareRepository
	customerService        *customerService.CustomerService
	teamService            *teamService.TeamService
}

func InitCustomerCareService(
	customerCareRepository *customerCareRepository.CustomerCareRepository,
	customerService *customerService.CustomerService,
	teamService *teamService.TeamService,
) *CustomerCareService {
	return &CustomerCareService{
		customerCareRepository: customerCareRepository,
		customerService:        customerService,
		teamService:            teamService,
	}
}

func (service *CustomerCareService) FindByID(id string) (customerCareModel.CustomerCare, error) {
	result, err := service.customerCareRepository.FindOneByField("ID", id, "Team", "Customer")
	customerCare, _ := result.(*customerCareModel.CustomerCare)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *customerCare, appException.NotFoundException(customerCareExceptionMessage.CUSTOMER_CARE_NOT_FOUND)
	}
	return *customerCare, nil
}

func (service *CustomerCareService) Create(
	dto customerCareDTO.CreateCustomerCareDTO,
	attendant attendantModel.Attendant,
) (customerCareModel.CustomerCare, error) {

	customerExists, err := service.customerService.FindByCPF(dto.CustomerCPF)
	if err != nil {
		return customerCareModel.CustomerCare{}, err
	}

	teamExists, err := service.teamService.FindByID(attendant.TeamID.String())
	if err != nil {
		return customerCareModel.CustomerCare{}, err
	}

	model := customerCareModel.CustomerCare{
		Date:        dto.Date,
		CustomerCPF: customerExists.CPF,
		TeamID:      attendant.TeamID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = service.customerCareRepository.Create(model)
	if err != nil {
		return customerCareModel.CustomerCare{}, err
	}

	model.Customer = customerExists
	model.Team = teamExists
	return model, err
}

func (service *CustomerCareService) FindAll(dto customerCareDTO.QueryCustomerCareDTO) (int64, []customerCareModel.CustomerCare, error) {
	var customerCares []customerCareModel.CustomerCare
	totalCount, err := service.customerCareRepository.FindAll(dto, &customerCares, "Team", "Customer")
	if err != nil {
		return 0, nil, err
	}
	return totalCount, customerCares, nil
}

func (service *CustomerCareService) Delete(id string) (customerCareModel.CustomerCare, error) {
	customerCareExists, err := service.FindByID(id)
	if err != nil {
		return customerCareModel.CustomerCare{}, err
	}

	err = service.customerCareRepository.Delete("ID", id)
	if err != nil {
		return customerCareModel.CustomerCare{}, err
	}
	return customerCareExists, nil
}