package appExceptions

import appDto "github.com/chronicler-org/core/src/app/dto"

type TNotFoundException struct {
	HttpException
}

func NotFoundException(errors []appDto.CustomErrorDTO) *TNotFoundException {
	return &TNotFoundException{HttpException: HttpException{Errors: errors, StatuCode: 404}}
}
