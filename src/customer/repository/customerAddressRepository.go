package customerRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerModel "github.com/chronicler-org/core/src/customer/model"
)

type CustomerAddressRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerAddressRepository(db *gorm.DB) *CustomerAddressRepository {
	return &CustomerAddressRepository{
		BaseRepository: appRepository.NewRepository(db, customerModel.CustomerAddress{}),
	}
}
