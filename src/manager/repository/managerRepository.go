package managerRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	managerModel "github.com/chronicler-org/core/src/manager/model"
)

type ManagerRepository struct {
	*appRepository.BaseRepository
}

func InitManagerRepository(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{
		BaseRepository: appRepository.NewRepository(db, managerModel.Manager{}),
	}
}
