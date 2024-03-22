package appException

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
)

type HttpException struct {
	Errors     []appDto.CustomErrorDTO
	StatusCode int
}

func (he *HttpException) getStatusCode() int {
	return he.StatusCode
}

func (he *HttpException) getErrors() []appDto.CustomErrorDTO {
	return he.Errors
}

func (he *HttpException) getErrorPagination() appUtil.PaginateErrorResponse {
	return appUtil.PaginateError(he.getErrors())
}

func (he *HttpException) Error() string {
	if len(he.Errors) > 0 {
		return he.Errors[0].Title
	}
	return ""
}
