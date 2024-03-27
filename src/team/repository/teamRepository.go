package tagRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	teamModel "github.com/chronicler-org/core/src/team/model"
)

type TeamRepository struct {
	*appRepository.BaseRepository
}

func InitTagRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		BaseRepository: appRepository.NewRepository(db, teamModel.Team{}),
	}
}
