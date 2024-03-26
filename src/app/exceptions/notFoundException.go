package appException

import (
	"net/http"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

func NotFoundException(error appDto.CustomErrorDTO) *HttpException {
	return &HttpException{
		Errors: []appDto.CustomErrorDTO{error},
		StatusCode: http.StatusNotFound,
	}
}
