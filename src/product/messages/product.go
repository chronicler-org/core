package productExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var PRODUCT_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "PRODUCT_NOT_FOUND",
	Title:  "Produto não encontrado",
	Detail: "Produto não foi encontrado na base de dados",
}

var OUT_OF_STOCK = appDto.CustomErrorDTO{
	Code:   "OUT_OF_STOCK",
	Title:  "Sem estoque suficiente",
	Detail: "O produto [ID: %s] não possui estoque suficiente. Estoque atual: %d. Quantidade solicitada: %d",
}
