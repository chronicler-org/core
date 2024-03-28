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
	customerRepository *customerRepository.CustomerRepository
	tagService         *tagService.TagService
}

func InitCustomerService(
	customerRepository *customerRepository.CustomerRepository,
	tagService *tagService.TagService,
) *CustomerService {
	return &CustomerService{
		customerRepository: customerRepository,
		tagService:         tagService,
	}
}

func (service *CustomerService) FindByID(id string) (customerModel.Customer, error) {
	result, err := service.customerRepository.FindByID(id, "Tags")
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
	var tags []*tagModel.Tag
	for _, tagID := range dto.TagIDs {
		tag, err := service.tagService.FindByID(tagID)
		if err != nil {
			return  customerModel.Customer{}, err
		}
		tags = append(tags, &tag)
	}
	model.Tags = tags;

	err := service.customerRepository.Create(model)
	if err != nil {
		return customerModel.Customer{}, err
	}

	return model, err
}

func (service *CustomerService) Update(id string, dto customerDTO.UpdateCustomerDTO) (customerModel.Customer, error) {
	customerExists, err := service.FindByID(id)
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

	customerUpdated, err := service.FindByID(id)

	return customerUpdated, err
}

func (service *CustomerService) FindAll(dto appDto.PaginationDTO) (int64, []customerModel.Customer, error) {
	var customers []customerModel.Customer
	totalCount, err := service.customerRepository.FindAll(dto.GetLimit(), dto.GetPage(), &customers, "Tags")
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

	err = service.customerRepository.Delete(id)
	if err != nil {
		return customerModel.Customer{}, err
	}
	return customerExists, nil
}

func (service *CustomerService) updateCustomerTags(customer *customerModel.Customer, tagIDs []string) error {
	err := service.customerRepository.ClearAssociations(customer.ID.String(), "Tags")
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

	err = service.customerRepository.ReplaceAssociations(customer.ID.String(), tags, "Tags")
	if err != nil {
		return err
	}
	customer.Tags = tags

	return nil
}
