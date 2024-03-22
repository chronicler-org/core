package appExceptions

import appDto "github.com/chronicler-org/core/src/app/dto"

type TUnauthorizedException struct {
	HttpException
}

func UnauthorizedException(errors []appDto.CustomErrorDTO) *TUnauthorizedException {
	return &TUnauthorizedException{HttpException: HttpException{Errors: errors, StatuCode: 401}}
}
