package managerExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var MANAGER_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "MANAGER_NOT_FOUND",
	Title:  "Gerente não encontrado",
	Detail: "Gerente não foi encontrado na base de dados",
}

var MANAGER_PASSWORDS_DONT_MATCH = appDto.CustomErrorDTO{
	Code:   "MANAGER_PASSWORD_DONT_MATCH",
	Title:  "As senhas informadas são diferentes",
	Detail: "As senhas informadas pelo gerente são diferentes.",
}
