package productRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	productModel "github.com/chronicler-org/core/src/product/model"
)

type ProductRepository struct {
	*appRepository.BaseRepository
}

func InitManagerRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: appRepository.NewRepository(db, productModel.Product{}),
	}
}
