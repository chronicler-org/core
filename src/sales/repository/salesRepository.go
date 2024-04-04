package salesRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	salesModel "github.com/chronicler-org/core/src/sales/model"
)

type SalesRepository struct {
	*appRepository.BaseRepository
}

func InitSalesRepository(db *gorm.DB) *SalesRepository {
	return &SalesRepository{
		BaseRepository: appRepository.NewRepository(db, salesModel.Sale{}),
	}
}
