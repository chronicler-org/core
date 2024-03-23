package appException

import (
	"net/http"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

func BadRequestException(errors []appDto.CustomErrorDTO) *HttpException {
	return &HttpException{
		Errors: errors, StatusCode: http.StatusBadRequest,
	}
}
