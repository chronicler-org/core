package salesRepository

import (
	"reflect"
	"time"

	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleEnum "github.com/chronicler-org/core/src/sales/enum"
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

func (r *SaleItemRepository) FindAll(dto interface{}, results interface{}, preloads ...string) (int64, error) {
	query := r.Db.Model(&salesModel.SaleItem{})
	query.Joins("INNER JOIN sales ON sales.customer_care_id = sale_items.sale_id")

	queryBuilder := appUtil.QueryBuilder(dto, query)
	queryBuilder.BuildQuery()
	queryBuilder.ApplyOrder()

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	offset, limit := queryBuilder.GetPagination()

	return count, query.Limit(limit).Offset(offset).Find(results).Error
}

func (r *SaleItemRepository) GetSaleProductSummary(
	dto salesDTO.QuerySalesProductSummaryDTO,
	results interface{},
) (int64, error) {
	query := r.Db.Model(&salesModel.SaleItem{})
	query.Joins("INNER JOIN sales ON sales.customer_care_id = sale_items.sale_id")

	queryBuilder := appUtil.QueryBuilder(dto, query)
	queryBuilder.BuildQuery()

	// Query to count total number of records
	var totalCount int64
	err := query.Group("product_id").Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	offset, limit := queryBuilder.GetPagination()
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

func (r *SaleItemRepository) GetTotalQuantitySoldByCreatedMonth(month time.Month, year int) (int64, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	var totalQuantity int64
	err := r.Db.Model(modelPtr).
		Select("COALESCE(SUM(sale_items.quantity), 0) as total_quantity").
		Joins("INNER JOIN sales ON sales.customer_care_id = sale_items.sale_id").
		Where("EXTRACT(MONTH FROM sale_items.created_at) = ?", month).
		Where("EXTRACT(YEAR FROM sale_items.created_at) = ?", year).
		Where("sales.status = ?", saleEnum.PURCHASE_COMPLETED).
		Pluck("total_quantity", &totalQuantity).Error

	if err != nil {
		return 0, err
	}

	return totalQuantity, nil
}
