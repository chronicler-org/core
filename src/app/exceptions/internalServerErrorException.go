package appException

import (
	"net/http"
	"strings"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type TInternalServerErrorException struct {
	HttpException
}

func InternalServerErrorException(message string) *TInternalServerErrorException {
	statusText := http.StatusText(http.StatusInternalServerError)
	code := strings.ReplaceAll(strings.ToUpper(statusText), " ", "_")

	return &TInternalServerErrorException{
		HttpException: HttpException{
			Errors: []appDto.CustomErrorDTO{{
				Code:   code,
				Title:  message,
				Detail: message,
			}},
			StatusCode: http.StatusInternalServerError,
		},
	}
}
