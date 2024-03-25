package managerRepository

import (
	appRepository "github.com/chronicler-org/core/src/app/repository"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	"gorm.io/gorm"
)

type ManagerRepository struct {
	*appRepository.BaseRepository
}

func InitManagerRepository(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{
		BaseRepository: appRepository.NewRepository(db, managerModel.Manager{}),
	}
}
