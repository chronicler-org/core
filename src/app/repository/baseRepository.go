package appRepository

import (
	"fmt"
	"reflect"

	appUtil "github.com/chronicler-org/core/src/app/utils"
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

func (r *BaseRepository) BeginTransaction() *gorm.DB {
	tx := r.Db.Begin()
	return tx
}

func (r *BaseRepository) Create(data interface{}) error {
	return r.Db.Model(&r.Model).Create(data).Error
}

func (r *BaseRepository) CreateWithTransaction(tx *gorm.DB, data interface{}) error {
	return tx.Model(r.Model).Create(data).Error
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

func (r *BaseRepository) UpdateWithTransaction(tx *gorm.DB, data interface{}) error {
	return tx.Save(data).Error
}

func (r *BaseRepository) FindAll(dto interface{}, results interface{}, preloads ...string) (int64, error) {
	queryBuilder := appUtil.QueryBuilder(dto, r.Db.Model(r.Model))
	query := queryBuilder.BuildQuery()

	var count int64
	err := query.Count(&count).Error
	if err != nil {
		return 0, err
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	offset, limit := queryBuilder.GetPagination()
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

func (r *BaseRepository) DeleteWithTransaction(tx *gorm.DB, field, value string) error {
	return tx.Delete(&r.Model, fmt.Sprintf("%s = ?", field), value).Error
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
