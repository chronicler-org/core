package appException

import (
	"net/http"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type TUnauthorizedException struct {
	HttpException
}

func UnauthorizedException(error appDto.CustomErrorDTO) *TUnauthorizedException {
	return &TUnauthorizedException{HttpException: HttpException{Errors: []appDto.CustomErrorDTO{error}, StatusCode: http.StatusUnauthorized}}
}
