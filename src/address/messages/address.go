package addressExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var ADDRESS_NOT_FOUND = appDto.CustomErrorDTO{
	Code: "ADDRESS_NOT_FOUND",
	Title: "Endereço não encontrado",
	Detail: "Endereço não foi encontrado na base de dados",
}