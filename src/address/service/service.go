package addressService

import (
	"time"

	addressDTO "github.com/chronicler-org/core/src/address/dto"
	addressModel "github.com/chronicler-org/core/src/address/model"
	addressRepository "github.com/chronicler-org/core/src/address/repository"
	"github.com/google/uuid"
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
	return service.repository.FindByID(id)
}

func (service *AddressService) Create(dto addressDTO.CreateAddressDTO) (uuid.UUID, error) {
	model := addressModel.Address{
		ID: uuid.New(),
		CEP: dto.CEP,
		City: dto.City,
		Number: dto.Number,
		Estate: dto.Estate,
		StreetName: dto.StreetName,
		Complement: dto.Complement,
		Neighborhood: dto.Neighborhood,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := service.repository.Create(model)

	return model.ID, err
}

func (service *AddressService) Update(id string, dto addressDTO.UpdateAddressDTO) (addressModel.Address, error) {
	updatedAddress, err := service.repository.FindByID(id)
	if err != nil {
		return updatedAddress, err
	}
	if updatedAddress.ID == uuid.Nil {
		return updatedAddress, err
	}
	if dto.CEP != "" {
		updatedAddress.CEP = dto.CEP
	}
	if dto.City != "" {
		updatedAddress.City = dto.City
	}
	if dto.Complement != "" {
		updatedAddress.Complement = dto.Complement
	}
	if dto.Estate != "" {
		updatedAddress.Estate = dto.Estate
	}
	if dto.Neighborhood != "" {
		updatedAddress.Neighborhood = dto.Neighborhood
	}
	if dto.Number != "" {
		updatedAddress.Number = dto.Number
	}
	if dto.StreetName != "" {
		updatedAddress.StreetName = dto.StreetName
	}
	updatedAddress.UpdatedAt = time.Now()

	err = service.repository.Update(updatedAddress)

	return updatedAddress, err
}