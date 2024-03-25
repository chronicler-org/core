package customerService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerExceptionMessage "github.com/chronicler-org/core/src/customer/messages"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
)

type CustomerService struct {
	repository *customerRepository.CustomerRepository
}

func InitCustomerService(r *customerRepository.CustomerRepository) *CustomerService {
	return &CustomerService{
		repository: r,
	}
}

func (service *CustomerService) FindByID(id string) (customerModel.Customer, error) {
	result, err := service.repository.FindByID(id)
	customer, _ := result.(*customerModel.Customer)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *customer, appException.NotFoundException(customerExceptionMessage.CUSTOMER_NOT_FOUND)
	}
	return *customer, nil
}

func (service *CustomerService) Create(dto customerDTO.CreateCustomerDTO) (customerModel.Customer, error) {
	model := customerModel.Customer{
		ID:        uuid.New(),
		CPF:       dto.CPF,
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Job:       dto.Job,
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := service.repository.Create(model)

	return model, err
}

func (service *CustomerService) Update(id string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	customerExists, err := service.FindByID(id)
	if err != nil {
		return customerModel.Customer{}, err
	}

	appUtil.UpdateModelFromDTO(&customerExists, dto)

	customerExists.UpdatedAt = time.Now()
	err = service.repository.Update(customerExists)
	return customerExists, err
}

func (service *CustomerService) FindAll(dto appDto.PaginationDTO) (int64, []customerModel.Customer, error) {
	var customers []customerModel.Customer
	totalCount, err := service.repository.FindAll(dto.GetLimit(), dto.GetPage(), &customers)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, customers, nil
}

func (service *CustomerService) Delete(id string) (customerModel.Customer, error) {
	customerExists, err := service.FindByID(id)
	if err != nil {
		return customerModel.Customer{}, err
	}

	err = service.repository.Delete(id)
	if err != nil {
		return customerModel.Customer{}, err
	}
	return customerExists, nil
}
