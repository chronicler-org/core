package salesRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	appUtil "github.com/chronicler-org/core/src/app/utils"
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

func (r *SaleItemRepository) GetSaleProductSummary(
	dto interface{},
	results interface{},
) (int64, error) {

	query, paginationDTO := appUtil.MapDTOToQuery(dto, r.Db.Model(&salesModel.SaleItem{}))

	// Query to count total number of records
	var totalCount int64
	err := query.Group("product_id").Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	page := paginationDTO.GetPage()
	limit := paginationDTO.GetLimit()
	offset := (page - 1) * limit

	err = query.
		Joins("JOIN products ON products.id = sale_items.product_id").
		Select("sale_items.product_id, products.model, SUM(sale_items.quantity) as total_quantity").
		Group("sale_items.product_id, products.model").
		Order("total_quantity DESC").
		Limit(limit).
		Offset(offset).
		Scan(results).Error
	if err != nil {
		return 0, err
	}

	return totalCount, err
}
