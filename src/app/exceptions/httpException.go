package exceptions

import (
	"github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/utils"
)

type HttpException struct {
	Errors []CustomErrorDTO
	Status int
}

func (he *HttpException) getStatus() int {
	return he.Status
}

func (he *HttpException) getErrors() []CustomErrorDTO {
	return he.Errors
}

func (he *HttpException) getPaginateResponse() PaginateErrorResponse {
	return PaginateError(he.getErrors())
}
