package appException

import appDto "github.com/chronicler-org/core/src/app/dto"

type TNotFoundException struct {
	HttpException
}

func NotFoundException(error appDto.CustomErrorDTO) *TNotFoundException {
	return &TNotFoundException{
		HttpException: HttpException{Errors: []appDto.CustomErrorDTO{error}, StatuCode: 404}}
}
