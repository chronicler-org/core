package appRepository

import (
	"fmt"
	"reflect"
	"time"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"gorm.io/gorm"
)

type BaseRepository struct {
	Db    *gorm.DB
	Model interface{}
}

func NewRepository(db *gorm.DB, model interface{}) *BaseRepository {
	return &BaseRepository{
		Db:    db,
		Model: model,
	}
}

func (r *BaseRepository) Create(data interface{}) error {
	return r.Db.Model(&r.Model).Create(data).Error
}

func (r *BaseRepository) FindOneByField(field string, value interface{}, preloads ...string) (interface{}, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	query := r.Db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	query = query.Where(fmt.Sprintf("%s = ?", field), value)
	err := query.First(modelPtr).Error
	return modelPtr, err
}

func (r *BaseRepository) Update(data interface{}) error {
	return r.Db.Save(data).Error
}

func (r *BaseRepository) FindAll(dto interface{}, results interface{}, preloads ...string) (int64, error) {
	query, paginationDTO := r.mapDTOToQuery(dto)

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	page := paginationDTO.GetPage()
	limit := paginationDTO.GetLimit()
	offset := (page - 1) * limit
	err = query.Limit(limit).Offset(offset).Find(results).Error
	return count, err
}

func (r *BaseRepository) Count() (int64, error) {
	var count int64
	err := r.Db.Model(&(r.Model)).Count(&count).Error
	return count, err
}

func (r *BaseRepository) Delete(field, value string) error {
	return r.Db.Delete(&r.Model, fmt.Sprintf("%s = ?", field), value).Error
}

func (r *BaseRepository) ReplaceAssociationsByField(field, value string, associations interface{}, associationName string) error {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	if err := r.Db.First(modelPtr, fmt.Sprintf("%s = ?", field), value).Error; err != nil {
		return err
	}

	if err := r.Db.Model(modelPtr).Association(associationName).Replace(associations); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) ClearAssociationsByField(field, value string, associationName string) error {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	if err := r.Db.First(modelPtr, fmt.Sprintf("%s = ?", field)).Error; err != nil {
		return err
	}

	if err := r.Db.Model(modelPtr).Association(associationName).Clear(); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) CountByCreatedMonth(month time.Month, year int) (int64, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	var count int64
	err := r.Db.Model(modelPtr).
		Where("EXTRACT(MONTH FROM created_at) = ?", month).
		Where("EXTRACT(YEAR FROM created_at) = ?", year).
		Count(&count).Error
	return count, err
}

func (r *BaseRepository) mapDTOToQuery(dto interface{}) (*gorm.DB, appDto.PaginationDTO) {
	var paginationDTO appDto.PaginationDTO

	dtoValue := reflect.ValueOf(dto)
	if dtoValue.Kind() == reflect.Ptr {
		dtoValue = dtoValue.Elem()
	}
	dtoType := dtoValue.Type()

	query := r.Db.Model(r.Model)
	for i := 0; i < dtoType.NumField(); i++ {
		fieldName := dtoType.Field(i).Name
		fieldValue := dtoValue.Field(i)

		switch fieldName {
		case "PaginationDTO":
			paginationDTOValue := fieldValue.Interface().(appDto.PaginationDTO)
			paginationDTO.Limit = paginationDTOValue.Limit
			paginationDTO.Page = paginationDTOValue.Page
		case "Limit":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int {
				paginationDTO.Limit = int(fieldValue.Int())
			}
		case "Page":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int {
				paginationDTO.Page = int(fieldValue.Int())
			}
		case "CreatedMonth":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int && fieldValue.Int() != 0 {
				query = query.Where("EXTRACT(MONTH FROM created_at) = ?", fieldValue.Int())
			}
		case "CreatedYear":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int && fieldValue.Int() != 0 {
				query = query.Where("EXTRACT(YEAR FROM created_at) = ?", int(fieldValue.Int()))
			}
		default:
			if fieldValue.Interface() != "" {
				tag := dtoType.Field(i).Tag.Get("query")
				if tag != "" {
					query = query.Where(fmt.Sprintf("%s = ?", tag), fieldValue.Interface())
				}
			}
		}

	}

	return query, paginationDTO
}
