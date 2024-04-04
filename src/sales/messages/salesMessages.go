package saleExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var SALE_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "SALE_NOT_FOUND",
	Title:  "Venda não encontrada",
	Detail: "Atendimento não foi encontrado na base de dados",
}

var INVALID_TRANSITION = appDto.CustomErrorDTO{
	Code:   "INVALID_TRANSITION",
	Title:  "Transição de status inválida",
	Detail: "A transição de status da venda solicitada é inválida",
}
