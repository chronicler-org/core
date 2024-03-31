package appUtil

import (
	"strings"
)

func GenerateEnumErrorDetail(fieldName string, validValues []interface{}) string {
	values := make([]string, len(validValues))
	for i, v := range validValues {
		values[i] = string(v.(string))
	}
	return "Os valores válidos para o campo " + fieldName + " são: " + strings.Join(values, ", ")
}
