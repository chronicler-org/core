package appUtil

import (
	"strings"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
)

func generateEnumErrorDetail(fieldName string, validValues []interface{}) string {
	values := make([]string, len(validValues))
	for i, v := range validValues {
		values[i] = string(v.(string))
	}
	return "Os valores válidos para o campo " + fieldName + " são: " + strings.Join(values, ", ")
}

func GenerateEnumInvalidFieldError(fieldName string, validValues []interface{}) error {
	customError := appDto.CustomErrorDTO{
		Code:   "INVALID_DATA",
		Title:  "Campo " + fieldName + " é inválido",
		Detail: generateEnumErrorDetail(fieldName, validValues),
	}

	return appException.BadRequestException([]appDto.CustomErrorDTO{customError})
}
