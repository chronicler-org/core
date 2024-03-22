package appException

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
)

type HttpException struct {
	Errors    []appDto.CustomErrorDTO
	StatuCode int
}

func (he *HttpException) getStatusCode() int {
	return he.StatuCode
}

func (he *HttpException) getErrors() []appDto.CustomErrorDTO {
	return he.Errors
}

func (he *HttpException) getErrorPagination() appUtil.PaginateErrorResponse {
	return appUtil.PaginateError(he.getErrors())
}
