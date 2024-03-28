package teanExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var TEAM_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "TEAM_NOT_FOUND",
	Title:  "Equipe não encontrado",
	Detail: "Equipe não foi encontrado na base de dados",
}