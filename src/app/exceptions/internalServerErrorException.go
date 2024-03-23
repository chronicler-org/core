package appException

import (
	"net/http"
	"strings"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

func InternalServerErrorException(message string) *HttpException {
	statusText := http.StatusText(http.StatusInternalServerError)
	code := strings.ReplaceAll(strings.ToUpper(statusText), " ", "_")

	return &HttpException{
		Errors: []appDto.CustomErrorDTO{{
			Code:   code,
			Title:  message,
		}},
		StatusCode: http.StatusInternalServerError,
	}
}
