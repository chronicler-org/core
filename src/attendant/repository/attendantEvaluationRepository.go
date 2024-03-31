package attendantRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
)

type AttendantEvaluationRepository struct {
	*appRepository.BaseRepository
}

func InitAttendantEvaluationRepository(db *gorm.DB) *AttendantEvaluationRepository {
	return &AttendantEvaluationRepository{
		BaseRepository: appRepository.NewRepository(db, attendantModel.AttendantEvaluation{}),
	}
}
