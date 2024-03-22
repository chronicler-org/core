package exceptions

import (
	"github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/exceptions"
)

type TNotFoundException struct {
	HttpException
}

func NotFoundException(errors []CustomErrorDTO) *TNotFoundException {
	return &TNotFoundException{HttpException: HttpException{Errors: errors, Status: 404}}
}
