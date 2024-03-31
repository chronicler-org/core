package authExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var LOGIN_FAILED = appDto.CustomErrorDTO{
	Code:   "LOGIN_FAILED",
	Title:  "Erro na realização de autenticação",
	Detail: "E-mail e/ou senha incorretos.",
}

var EXPIRED_AT = appDto.CustomErrorDTO{
	Code:   "EXPIRED_TOKEN",
	Title:  "Access Token expirado",
	Detail: "Token informado já expirou. Realize nova autenticação para adquirir novo token.",
}

var INVALID_AT = appDto.CustomErrorDTO{
	Code:   "INVALID_TOKEN",
	Title:  "Access Token inválido",
	Detail: "Token informado é inválido. Realize nova autenticação para adquirir novo token.",
}

var USER_SERVICE_ACCESS_INFO_DENIED = appDto.CustomErrorDTO{
	Code:   "USER_SERVICE_ACCESS_INFO_DENIED",
	Title:  "Operação não permitida",
	Detail: "Usuário ou seviço não tem permissão para acessar este recurso.",
}
