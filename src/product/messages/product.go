package productExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var PRODUCT_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "PRODUCT_NOT_FOUND",
	Title:  "Produto não encontrado",
	Detail: "Produto não foi encontrado na base de dados",
}