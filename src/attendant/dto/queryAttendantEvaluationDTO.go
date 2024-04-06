package attendantDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryAttendantEvaluationDTO struct {
	appDto.PaginationDTO
	AvaluatedID string `validate:"omitempty,uuid" query:"avaluated_id" pg:"operation=equal"`
	AvaluatorID string `validate:"omitempty,uuid" query:"avaluator_id" pg:"operation=equal"`
}
