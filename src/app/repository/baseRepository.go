package appRepository

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	db    *gorm.DB
	model interface{}
}

func NewRepository(db *gorm.DB, model interface{}) *BaseRepository {
	return &BaseRepository{
		db:    db,
		model: model,
	}
}

func (r *BaseRepository) Create(data interface{}) error {
	return r.db.Model(&r.model).Create(data).Error
}

func (r *BaseRepository) FindByID(id interface{}) (interface{}, error) {
	var result interface{}
	err := r.db.Where("id = ?", id).First(&result).Error
	return result, err
}

func (r *BaseRepository) Update(data interface{}) error {
	return r.db.Save(data).Error
}

func (r *BaseRepository) FindAll(limit, page int, results interface{}) (int64, error) {
	var count int64
	offset := (page - 1) * limit
	err := r.db.Model(r.model).Count(&count).Error
	if err != nil {
			return 0, err
	}
	err = r.db.Limit(limit).Offset(offset).Find(results).Error
	return count, err
}

func (r *BaseRepository) Count() (int64, error) {
	var count int64
	err := r.db.Model(&(r.model)).Count(&count).Error
	return count, err
}

func (r *BaseRepository) Delete(id string) error {
	return r.db.Delete(&r.model, "id = ?", id).Error
}
