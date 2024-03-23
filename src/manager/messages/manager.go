package managerExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var MANAGER_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "MANAGER_NOT_FOUND",
	Title:  "Gerente não encontrado",
}