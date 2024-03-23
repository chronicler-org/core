package appException

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type HttpException struct {
	Errors     []appDto.CustomErrorDTO
	StatusCode int
}

func (he *HttpException) GetStatusCode() int {
	return he.StatusCode
}

func (he *HttpException) GetErrors() []appDto.CustomErrorDTO {
	return he.Errors
}

func (he *HttpException) Error() string {
	if len(he.Errors) > 0 {
		return he.Errors[0].Code
	}
	return ""
}
