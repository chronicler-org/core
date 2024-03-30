package customerCareRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
)

type CustomerCareRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerCareRepository(db *gorm.DB) *CustomerCareRepository {
	return &CustomerCareRepository{
		BaseRepository: appRepository.NewRepository(db, customerCareModel.CustomerCare{}),
	}
}
