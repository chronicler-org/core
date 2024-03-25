package attendantRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
)

type AttendantRepository struct {
	*appRepository.BaseRepository
}

func InitAttendantRepository(db *gorm.DB) *AttendantRepository {
	return &AttendantRepository{
		BaseRepository: appRepository.NewRepository(db, attendantModel.Attendant{}),
	}
}
