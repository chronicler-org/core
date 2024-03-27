package appRepository

import (
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BaseRepository struct {
	Db    *gorm.DB
	model interface{}
}

func NewRepository(db *gorm.DB, model interface{}) *BaseRepository {
	return &BaseRepository{
		Db:    db,
		model: model,
	}
}

func (r *BaseRepository) Create(data interface{}) error {
	return r.Db.Model(&r.model).Create(data).Error
}

func (r *BaseRepository) FindByID(id interface{}, preloads ...string) (interface{}, error) {
	modelType := reflect.TypeOf(r.model)
	modelPtr := reflect.New(modelType).Interface()

	query := r.Db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	err := query.Where("id = ?", id).First(modelPtr).Error
	return modelPtr, err
}

func (r *BaseRepository) Update(data interface{}) error {
	return r.Db.Save(data).Error
}

func (r *BaseRepository) FindAll(limit, page int, results interface{}, preloads ...string) (int64, error) {
	var count int64
	offset := (page - 1) * limit
	err := r.Db.Model(r.model).Count(&count).Error
	if err != nil {
		return 0, err
	}
	err = r.Db.Preload(clause.Associations).Limit(limit).Offset(offset).Find(results).Error
	return count, err
}

func (r *BaseRepository) Count() (int64, error) {
	var count int64
	err := r.Db.Model(&(r.model)).Count(&count).Error
	return count, err
}

func (r *BaseRepository) Delete(id string) error {
	return r.Db.Delete(&r.model, "id = ?", id).Error
}

func (r *BaseRepository) ReplaceAssociations(modelID string, associations interface{}, associationName string) error {
	modelType := reflect.TypeOf(r.model)
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
	modelType := reflect.TypeOf(r.model)
	modelPtr := reflect.New(modelType).Interface()

	if err := r.Db.First(modelPtr, "id = ?", modelID).Error; err != nil {
		return err
	}

	if err := r.Db.Model(modelPtr).Association(associationName).Clear(); err != nil {
		return err
	}

	return nil
}
