package salesRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	salesModel "github.com/chronicler-org/core/src/sales/model"
)

type SaleRepository struct {
	*appRepository.BaseRepository
}

func InitSaleRepository(db *gorm.DB) *SaleRepository {
	return &SaleRepository{
		BaseRepository: appRepository.NewRepository(db, salesModel.Sale{}),
	}
}
