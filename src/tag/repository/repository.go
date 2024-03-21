package tagRepository

import (
	tagModel "github.com/chronicler-org/core/src/tag/model"
	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func InitTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

func (repository *TagRepository) Create(tag tagModel.Tag) error {
	return repository.db.Model(&tagModel.Tag{}).Create(tag).Error
}

func (repository *TagRepository) FindByID(id string) (tagModel.Tag, error) {
	var tag tagModel.Tag
	err := repository.db.Find(&tag, "id = ?", id).Error
	return tag, err
}

func (repository *TagRepository) Update(updatedTag tagModel.Tag) error {
	return repository.db.Save(updatedTag).Error
}

func (repository *TagRepository) FindAll() ([]tagModel.Tag, error) {
	var tags []tagModel.Tag
	err := repository.db.Find(&tags).Error
	return tags, err
}

func (repository *TagRepository) Delete(id string) error {
	return repository.db.Delete(&tagModel.Tag{}, "id = ?", id).Error
}
