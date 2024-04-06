package salesRepository

import (
	"time"

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

	applyDateFilter := dto.Month != 0 || dto.Year != 0
	var (
		startTime, endTime time.Time
	)

	// Set the time period to the specified month and year if provided
	if applyDateFilter {
		switch {
		case dto.Month != 0 && dto.Year != 0:
			startTime = time.Date(dto.Year, time.Month(dto.Month), 1, 0, 0, 0, 0, time.UTC)
			endTime = startTime.AddDate(0, 1, 0).Add(-time.Nanosecond)
		case dto.Month != 0:
			year, _, _ := time.Now().Date()
			startTime = time.Date(year, time.Month(dto.Month), 1, 0, 0, 0, 0, time.UTC)
			endTime = time.Date(year, time.Month(dto.Month)+1, 1, 0, 0, 0, -1, time.UTC)
		case dto.Year != 0:
			startTime = time.Date(dto.Year, time.January, 1, 0, 0, 0, 0, time.UTC)
			endTime = time.Date(dto.Year+1, time.January, 1, 0, 0, 0, -1, time.UTC)
		}
	}

	query = query.
		Select("DATE(created_at) as sale_date, SUM(total_value) as total_value")
	queryBuilder.BuildQuery()
	query = query.Group("DATE(created_at)")

	if applyDateFilter {
		query = query.Where("created_at BETWEEN ? AND ?", startTime, endTime)
	}

	// Query to count total number of records
	var totalCount int64
	err := query.Count(&totalCount).Error
	if err != nil {
		return 0, err
	}

	offset, limit := queryBuilder.GetPagination()
	queryBuilder.ApplyOrder()
	err = query.
		Limit(limit).
		Offset(offset).
		Scan(results).Error

	if err != nil {
		return 0, err
	}

	return totalCount, nil
}
