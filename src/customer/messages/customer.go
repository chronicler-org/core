package customerExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var CUSTOMER_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "CUSTOMER_NOT_FOUND",
	Title:  "Cliente não encontrado",
	Detail: "Cliente não foi encontrado na base de dados",
}