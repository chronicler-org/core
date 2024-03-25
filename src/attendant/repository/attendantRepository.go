package attendantRepository

import (
	appRepository "github.com/chronicler-org/core/src/app/repository"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	"gorm.io/gorm"
)

type AttendantRepository struct {
	*appRepository.BaseRepository
}

func InitAttendantRepository(db *gorm.DB) *AttendantRepository {
	return &AttendantRepository{
		BaseRepository: appRepository.NewRepository(db, attendantModel.Attendant{}),
	}
}
