package appException

import appDto "github.com/chronicler-org/core/src/app/dto"

type TForbiddenException struct {
	HttpException
}

func ForbiddenException(error appDto.CustomErrorDTO) *TForbiddenException {
	return &TForbiddenException{HttpException: HttpException{Errors: []appDto.CustomErrorDTO{error}, StatuCode: 403}}
}
