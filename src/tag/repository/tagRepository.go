package tagRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	tagModel "github.com/chronicler-org/core/src/tag/model"
)

type TagRepository struct {
	*appRepository.BaseRepository
}

func InitTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		BaseRepository: appRepository.NewRepository(db, tagModel.Tag{}),
	}
}
