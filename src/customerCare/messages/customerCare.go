package customerCareExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var CUSTOMER_CARE_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "CUSTOMER_CARE_NOT_FOUND",
	Title:  "Atendimento não encontrado",
	Detail: "Atendimento não foi encontrado na base de dados",
}

var CUSTOMER_CARE_EVALUATION_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "CUSTOMER_CARE_EVALUATION_NOT_FOUND",
	Title:  "Avaliação de atendimento não encontrada",
	Detail: "O atendimento não possui avaliação.",
}

var CUSTOMER_CARE_ALREADY_EVALUATED = appDto.CustomErrorDTO{
	Code:   "CUSTOMER_CARE_ALREADY_EVALUATED",
	Title:  "Avaliação de atendimento já cadastrada",
	Detail: "O atendimento de cliente já possui uma avaliação cadastrada.",
}