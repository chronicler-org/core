package customerRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerModel "github.com/chronicler-org/core/src/customer/model"
)

type CustomerRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		BaseRepository: appRepository.NewRepository(db, customerModel.Customer{}),
	}
}
