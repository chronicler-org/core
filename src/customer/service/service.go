package customerService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerExceptionMessage "github.com/chronicler-org/core/src/customer/messages"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
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

func (service *CustomerService) Create(dto customerDTO.CreateCustomerDTO) (uuid.UUID, error) {
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

	return model.ID, err
}

func (service *CustomerService) Update(id string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	// implementar validacao de dados
	updatedCustomer, err := service.repository.FindByID(id)
	if err != nil {
		return updatedCustomer, err
	}
	if updatedCustomer.ID == uuid.Nil {
		return updatedCustomer, serviceErrors.NewError("Cliente não encontrado")
	}
	if dto.Name != "" {
		updatedCustomer.Name = dto.Name
	}
	if dto.Email != "" {
		updatedCustomer.Email = dto.Email
	}
	if dto.Phone != "" {
		updatedCustomer.Phone = dto.Phone
	}
	if dto.Job != "" {
		updatedCustomer.Job = dto.Job
	}
	if !dto.BirthDate.IsZero() {
		updatedCustomer.BirthDate = dto.BirthDate
	}
	updatedCustomer.UpdatedAt = time.Now()

	err = service.repository.Update(updatedCustomer)

	return updatedCustomer, err
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
