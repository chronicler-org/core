package managerRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerServiceModel "github.com/chronicler-org/core/src/customerService/model"
)

type CustomerServiceRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerServiceRepository(db *gorm.DB) *CustomerServiceRepository {
	return &CustomerServiceRepository{
		BaseRepository: appRepository.NewRepository(db, customerServiceModel.CustomerService{}),
	}
}
