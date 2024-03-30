package customerCareExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var SERVICE_EVALUATION_CARE_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "SERVICE_EVALUATION_CARE_NOT_FOUND",
	Title:  "Avaliação de atendimento não encontrada",
	Detail: "Avaliação de atendimento não foi encontrada na base de dados",
}
