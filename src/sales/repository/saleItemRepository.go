package salesRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	salesModel "github.com/chronicler-org/core/src/sales/model"
)

type SaleItemRepository struct {
	*appRepository.BaseRepository
}

func InitSaleItemRepository(db *gorm.DB) *SaleItemRepository {
	return &SaleItemRepository{
		BaseRepository: appRepository.NewRepository(db, salesModel.SaleItem{}),
	}
}
 