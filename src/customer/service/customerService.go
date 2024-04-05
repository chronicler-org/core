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
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

type CustomerService struct {
	customerRepository        *customerRepository.CustomerRepository
	customerAddressRepository *customerRepository.CustomerAddressRepository
	tagService                *tagService.TagService
}

func InitCustomerService(
	customerRepository *customerRepository.CustomerRepository,
	customerAddressRepository *customerRepository.CustomerAddressRepository,
	tagService *tagService.TagService,
) *CustomerService {
	return &CustomerService{
		customerRepository:        customerRepository,
		customerAddressRepository: customerAddressRepository,
		tagService:                tagService,
	}
}

func (service *CustomerService) FindCustomerByCPF(cpf string) (customerModel.Customer, error) {
	result, err := service.customerRepository.FindOneByField("CPF", cpf, "Tags", "Address")
	customer, _ := result.(*customerModel.Customer)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *customer, appException.NotFoundException(customerExceptionMessage.CUSTOMER_NOT_FOUND)
	}
	return *customer, nil
}

func (service *CustomerService) CreateCustomer(dto customerDTO.CreateCustomerDTO) (customerModel.Customer, error) {
	customerAddressExists, err := service.FindCustomerAddressByID(dto.AddressId)
	if err != nil {
		return customerModel.Customer{}, err
	}

	model := customerModel.Customer{
		CPF:       dto.CPF,
		Name:      dto.Name,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Job:       dto.Job,
		AddressID: customerAddressExists.ID,
		BirthDate: dto.BirthDate,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var tags []*tagModel.Tag
	for _, tagID := range dto.TagIDs {
		tag, err := service.tagService.FindByID(tagID)
		if err != nil {
			return customerModel.Customer{}, err
		}
		tags = append(tags, &tag)
	}
	model.Tags = tags

	err = service.customerRepository.Create(model)
	if err != nil {
		return customerModel.Customer{}, err
	}
	model.Address = customerAddressExists
	return model, err
}

func (service *CustomerService) UpdateCustomer(cpf string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	customerExists, err := service.FindCustomerByCPF(cpf)
	if err != nil {
		return customerModel.Customer{}, err
	}
	customerAddress := customerModel.CustomerAddress{}

	if dto.AddressId != "" {
		customerAddressExists, err := service.FindCustomerAddressByID(dto.AddressId)
		if err != nil {
			return customerModel.Customer{}, err
		}
		customerAddress = customerAddressExists
	}

	if err != nil {
		return customerModel.Customer{}, err
	}

	appUtil.UpdateModelFromDTO(&customerExists, &dto)

	err = service.updateCustomerTags(&customerExists, dto.TagIDs)
	if err != nil {
		return customerModel.Customer{}, err
	}

	customerExists.UpdatedAt = time.Now()
	err = service.customerRepository.Update(customerExists)
	if err != nil {
		return customerModel.Customer{}, err
	}

	customerUpdated, err := service.FindCustomerByCPF(cpf)
	customerUpdated.Address = customerAddress
	return customerUpdated, err
}

func (service *CustomerService) FindAllCustomers(dto customerDTO.CustomerQueryDTO) (int64, []customerModel.Customer, error) {
	var customers []customerModel.Customer
	totalCount, err := service.customerRepository.FindAll(dto, &customers, "Tags", "Address")
	if err != nil {
		return 0, nil, err
	}
	return totalCount, customers, nil
}

func (service *CustomerService) DeleteCustomer(cpf string) (customerModel.Customer, error) {
	customerExists, err := service.FindCustomerByCPF(cpf)
	if err != nil {
		return customerModel.Customer{}, err
	}

	err = service.customerRepository.Delete("CPF", cpf)
	if err != nil {
		return customerModel.Customer{}, err
	}
	return customerExists, nil
}

func (service *CustomerService) GetNewCustomersVariationPercent() (any, error) {
	type NewCustomersVariationDTO struct {
		PercentVariation float64 `json:"percent_variation"`
	}

	currentMonth := time.Now().Month()
	currentYear := time.Now().Year()
	currentMonthCount, err := service.customerRepository.CountCustomersByCreatedMonth(currentMonth, currentYear)

	if err != nil {
		return NewCustomersVariationDTO{}, err
	}

	lastMonth, lastYear := appUtil.GetLastMonth()
	lastMonthCount, err := service.customerRepository.CountCustomersByCreatedMonth(lastMonth, lastYear)
	if err != nil {
		return NewCustomersVariationDTO{}, err
	}

	percentVariation := appUtil.CalculatePercentVariation(float64(currentMonthCount), float64(lastMonthCount))

	return NewCustomersVariationDTO{
		PercentVariation: percentVariation,
	}, nil
}

func (service *CustomerService) updateCustomerTags(customer *customerModel.Customer, tagIDs []string) error {
	err := service.customerRepository.ClearAssociationsByField("CPF", customer.CPF, "Tags")
	if err != nil {
		return err
	}
	var tags []*tagModel.Tag
	for _, tagID := range tagIDs {
		tag, err := service.tagService.FindByID(tagID)
		if err != nil {
			return err
		}
		tags = append(tags, &tag)
	}

	err = service.customerRepository.ReplaceAssociationsByField("CPF", customer.CPF, tags, "Tags")
	if err != nil {
		return err
	}
	customer.Tags = tags

	return nil
}

func (service *CustomerService) FindCustomerAddressByID(id string) (customerModel.CustomerAddress, error) {
	result, err := service.customerAddressRepository.FindOneByField("ID", id)
	customerAddress, _ := result.(*customerModel.CustomerAddress)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *customerAddress, appException.NotFoundException(customerExceptionMessage.ADDRESS_NOT_FOUND)
	}
	return *customerAddress, nil
}

func (service *CustomerService) CreateCustomerAddress(dto customerDTO.CreateCustomerAddressDTO) (customerModel.CustomerAddress, error) {
	model := customerModel.CustomerAddress{
		ID:           uuid.New(),
		CEP:          dto.CEP,
		City:         dto.City,
		Number:       dto.Number,
		Estate:       dto.Estate,
		StreetName:   dto.StreetName,
		Complement:   dto.Complement,
		Neighborhood: dto.Neighborhood,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := service.customerAddressRepository.Create(model)

	return model, err
}

func (service *CustomerService) UpdateCustomerAddress(id string, dto customerDTO.UpdateCustomerAddressDTO) (customerModel.CustomerAddress, error) {
	customerAddressExists, err := service.FindCustomerAddressByID(id)
	if err != nil {
		return customerModel.CustomerAddress{}, err
	}

	appUtil.UpdateModelFromDTO(&customerAddressExists, &dto)

	customerAddressExists.UpdatedAt = time.Now()
	err = service.customerAddressRepository.Update(customerAddressExists)
	return customerAddressExists, err
}

func (service *CustomerService) FindAllCustomerAddresses(dto appDto.PaginationDTO) (int64, []customerModel.CustomerAddress, error) {
	var customerAddresses []customerModel.CustomerAddress
	totalCount, err := service.customerAddressRepository.FindAll(dto, &customerAddresses)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, customerAddresses, nil
}

func (service *CustomerService) DeleteCustomerAddress(id string) (customerModel.CustomerAddress, error) {
	customerAddressExists, err := service.FindCustomerAddressByID(id)
	if err != nil {
		return customerModel.CustomerAddress{}, err
	}

	err = service.customerAddressRepository.Delete("ID", id)
	if err != nil {
		return customerModel.CustomerAddress{}, err
	}

	return customerAddressExists, nil
}
