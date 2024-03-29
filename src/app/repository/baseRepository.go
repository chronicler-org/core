package appRepository

import (
	"fmt"
	"reflect"

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

func (r *BaseRepository) FindByID(id interface{}, preloads ...string) (interface{}, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	query := r.Db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.Where("id = ?", id).First(modelPtr).Error
	return modelPtr, err
}

func (r *BaseRepository) FindOneBy(by, value interface{}, preloads ...string) (interface{}, error) {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	query := r.Db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.Where(fmt.Sprintf("%s = ?", by), value).First(modelPtr).Error
	return modelPtr, err
}

func (r *BaseRepository) Update(data interface{}) error {
	return r.Db.Save(data).Error
}

func (r *BaseRepository) FindAll(dto interface{}, results interface{}, preloads ...string) (int64, error) {
	var count int64
	var paginationDTO appDto.PaginationDTO

	dtoValue := reflect.ValueOf(dto)
	if dtoValue.Kind() == reflect.Ptr {
		dtoValue = dtoValue.Elem()
	}
	dtoType := dtoValue.Type()

	filters := make(map[string]interface{})

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
		default:
			if fieldValue.Interface() != "" {
				tag := dtoType.Field(i).Tag.Get("query")
				if tag != "" {
					filters[tag] = fieldValue.Interface()
				}
			}
		}

	}
	query := r.Db.Model(r.Model)
	for key, value := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}

	err := query.Where(filters).Count(&count).Error
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
	if field != "" {
		return r.Db.Delete(&r.Model, fmt.Sprintf("%s = ?", field), value).Error
	}
	return r.Db.Delete(&r.Model, "id = ?", value).Error
}

func (r *BaseRepository) ReplaceAssociations(modelID string, associations interface{}, associationName string) error {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	if err := r.Db.First(modelPtr, "id = ?", modelID).Error; err != nil {
		return err
	}

	if err := r.Db.Model(modelPtr).Association(associationName).Replace(associations); err != nil {
		return err
	}

	return nil
}

func (r *BaseRepository) ClearAssociations(modelID string, associationName string) error {
	modelType := reflect.TypeOf(r.Model)
	modelPtr := reflect.New(modelType).Interface()

	if err := r.Db.First(modelPtr, "id = ?", modelID).Error; err != nil {
		return err
	}

	if err := r.Db.Model(modelPtr).Association(associationName).Clear(); err != nil {
		return err
	}

	return nil
}
