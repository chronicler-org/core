package customerCareRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
)

type ServiceEvaluationRepository struct {
	*appRepository.BaseRepository
}

func InitServiceEvaluationRepository(db *gorm.DB) *ServiceEvaluationRepository {
	return &ServiceEvaluationRepository{
		BaseRepository: appRepository.NewRepository(db, customerCareModel.ServiceEvaluation{}),
	}
}
