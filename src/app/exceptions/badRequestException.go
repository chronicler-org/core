package appException

import (
	"net/http"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type TBadRequestException struct {
	HttpException
}

func BadRequestException(errors []appDto.CustomErrorDTO) *TBadRequestException {
	return &TBadRequestException{HttpException: HttpException{Errors: errors, StatusCode: http.StatusBadRequest}}
}
