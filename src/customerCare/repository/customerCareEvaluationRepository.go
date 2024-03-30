package customerCareRepository

import (
	"gorm.io/gorm"

	appRepository "github.com/chronicler-org/core/src/app/repository"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
)

type CustomerCareEvaluationRepository struct {
	*appRepository.BaseRepository
}

func InitCustomerCareEvaluationRepository(db *gorm.DB) *CustomerCareEvaluationRepository {
	return &CustomerCareEvaluationRepository{
		BaseRepository: appRepository.NewRepository(db, customerCareModel.CustomerCareEvaluation{}),
	}
}
