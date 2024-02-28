package customerService

import (
	"time"

	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	return service.repository.FindByID(id)
}

func (service *CustomerService) Create(dto customerDTO.CreateCustomerDTO) (uuid.UUID, error) {
	newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return uuid.Nil, err
	}

	model := customerModel.Customer{
		ID:        uuid.New(),
		CPF:       dto.CPF,
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Job:       dto.Job,
		Password:  string(newPassword),
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = service.repository.Create(model)

	return model.ID, err
}

func (service *CustomerService) Update(id string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	updatedCustomer, err := service.repository.FindByID(id)
	if err != nil {
		return updatedCustomer, err
	}
	if updatedCustomer.ID == uuid.Nil {
		return updatedCustomer, err
	}

	if dto.CPF != "" {
		updatedCustomer.CPF = dto.CPF
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
	if dto.Password != "" {
		newPassword, err := bcrypt.GenerateFromPassword([]byte(dto.CPF), 10)
		if err != nil {
			return customerModel.Customer{}, err
		}
		updatedCustomer.Password = string(newPassword)
	}
	if !dto.BirthDate.IsZero() {
		updatedCustomer.BirthDate = dto.BirthDate
	}
	updatedCustomer.UpdatedAt = time.Now()

	err = service.repository.Update(updatedCustomer)

	return updatedCustomer, err
}

func (service *CustomerService) FindAll() ([]customerModel.Customer, error) {
	return service.repository.FindAll()
}

func (service *CustomerService) Delete(id string) error {
	return service.repository.Delete(id)
}
