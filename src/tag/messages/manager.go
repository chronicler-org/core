package tagExceptionMessage

import appDto "github.com/chronicler-org/core/src/app/dto"

var TAG_NOT_FOUND = appDto.CustomErrorDTO{
	Code:   "TAG_NOT_FOUND",
	Title:  "Tag não encontrado",
	Detail: "Tag não foi encontrado na base de dados",
}