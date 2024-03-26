package addressService

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	addressDTO "github.com/chronicler-org/core/src/address/dto"
	addressExceptionMessage "github.com/chronicler-org/core/src/address/messages"
	addressModel "github.com/chronicler-org/core/src/address/model"
	addressRepository "github.com/chronicler-org/core/src/address/repository"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
)

type AddressService struct {
	repository *addressRepository.AddressRepository
}

func InitAddressService(r *addressRepository.AddressRepository) *AddressService {
	return &AddressService{
		repository: r,
	}
}

func (service *AddressService) FindByID(id string) (addressModel.Address, error) {
	result, err := service.repository.FindByID(id)
	address, _ := result.(*addressModel.Address)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *address, appException.NotFoundException(addressExceptionMessage.ADDRESS_NOT_FOUND)
	}
	return *address, nil
}

func (service *AddressService) Create(dto addressDTO.CreateAddressDTO) (uuid.UUID, error) {
	model := addressModel.Address{
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

	err := service.repository.Create(model)

	return model.ID, err
}

func (service *AddressService) Update(id string, dto addressDTO.UpdateAddressDTO) (addressModel.Address, error) {
	addressExists, err := service.FindByID(id)
	if err != nil {
		return addressModel.Address{}, err
	}

	appUtil.UpdateModelFromDTO(&addressExists, dto)

	addressExists.UpdatedAt = time.Now()
	err = service.repository.Update(addressExists)
	return addressExists, err
}

func (service *AddressService) Delete(id string) (addressModel.Address, error) {
	addressExists, err := service.FindByID(id)
	if err != nil {
		return addressModel.Address{}, err
	}

	err = service.repository.Delete(id)
	if err != nil {
		return addressModel.Address{}, err
	}

	return addressExists, nil
}
