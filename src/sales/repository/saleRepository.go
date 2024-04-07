package salesRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
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

func (r *SaleItemRepository) GetTotalValuesSold(
	dto salesDTO.QueryTotalSalesSoldDTO,
	results interface{},
) (int64, error) {
	query := r.Db.Model(&salesModel.Sale{})
	queryBuilder := appUtil.QueryBuilder(dto, query)
	queryBuilder.BuildQuery()

	startDateStr := dto.StartDate.Format("2006-01-02 15:04:05")
	endDateStr := dto.EndDate.Format("2006-01-02 15:04:05")

	query = query.
		Select(`
			date_series::date AS sale_date,
			COALESCE(SUM(CASE WHEN sales.status = 'Compra concluida' THEN total_value ELSE 0 END), 0) AS total_value
		`).
		Joins("RIGHT JOIN generate_series(?::timestamp, ?::timestamp, '1 day') AS date_series ON DATE(sales.created_at) = date_series", startDateStr, endDateStr).
		Group("date_series::date").
		Order("date_series::date")

	// Query to count total number of records
	var totalCount int64
	err := query.Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	offset, limit := queryBuilder.GetPagination()
	err = query.
		Limit(limit).
		Offset(offset).
		Scan(results).Error

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
