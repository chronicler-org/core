package attendantExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var ATTENDANT_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "ATTENDANT_NOT_FOUND",
	Title:  "Atendente não encontrado",
	Detail: "Atendente não foi encontrado na base de dados",
}

var ATTENDANT_EVALUATION_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "ATTENDANT_EVALUATION_NOT_FOUND",
	Title:  "Avaliação do atendente não encontrado",
	Detail: "Avaliação do atendente não foi encontrado na base de dados",
}

var ATTENDANT_EVALUATION_SELF_EVALUATION = appDto.CustomErrorDTO{
	Code:   "ATTENDANT_EVALUATION_SELF_EVALUATION",
	Title:  "Avaliação do Atendente Inválida",
	Detail: "Um atendente não pode avaliar a si mesmo.",
}