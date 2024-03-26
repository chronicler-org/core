package addressRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	addressModel "github.com/chronicler-org/core/src/address/model"
)

type AddressRepository struct {
	*appRepository.BaseRepository
}

func InitAddressRepository(db *gorm.DB) *AddressRepository {
	return &AddressRepository{
		BaseRepository: appRepository.NewRepository(db, addressModel.Address{}),
	}
}