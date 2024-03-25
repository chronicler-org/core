package attendantExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var ATTENDANT_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "ATTENDANT_NOT_FOUND",
	Title:  "Atendente não encontrado",
	Detail: "Atendente não foi encontrado na base de dados",
}