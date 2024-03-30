package authExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var LOGIN_FAILED = appDto.CustomErrorDTO{
	Code:   "LOGIN_FAILED",
	Title:  "Erro na realização de autenticação",
	Detail: "E-mail e/ou senha incorretos.",
}
