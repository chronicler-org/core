package customerRepository

import (
	"reflect"
	"time"

	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerModel "github.com/chronicler-org/core/src/customer/model"
)

type CustomerRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerRepository(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		BaseRepository: appRepository.NewRepository(db, customerModel.Customer{}),
	}
}

func (r *CustomerRepository) CountCustomersByCreatedMonth(month time.Month, year int) (int64, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	var count int64
	err := r.Db.Model(modelPtr).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Count(&count).Error
	return count, err
}
