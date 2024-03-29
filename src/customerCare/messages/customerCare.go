package customerCareExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var CUSTOMER_CARE_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "CUSTOMER_CARE_NOT_FOUND",
	Title:  "Atendimento não encontrado",
	Detail: "Atendimento não foi encontrado na base de dados",
}