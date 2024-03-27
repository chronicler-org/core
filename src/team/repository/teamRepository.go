package teamRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	teamModel "github.com/chronicler-org/core/src/team/model"
)

type TeamRepository struct {
	*appRepository.BaseRepository
}

func InitTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		BaseRepository: appRepository.NewRepository(db, teamModel.Team{}),
	}
}
