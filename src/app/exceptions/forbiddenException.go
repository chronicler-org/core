package appExceptions

import appDto "github.com/chronicler-org/core/src/app/dto"

type TForbiddenException struct {
	HttpException
}

func ForbiddenException(errors []appDto.CustomErrorDTO) *TForbiddenException {
	return &TForbiddenException{HttpException: HttpException{Errors: errors, StatuCode: 403}}
}
