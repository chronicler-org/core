package appUtil

import (
	"reflect"
)

func UpdateModelFromDTO(model interface{}, dto interface{}) {
	dtoType := reflect.TypeOf(dto)
	dtoValue := reflect.ValueOf(dto)

	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)

	if modelType.Kind() != reflect.Ptr || modelType.Elem().Kind() != reflect.Struct {
		panic("O modelo fornecido deve ser um ponteiro para uma estrutura")
	}

	modelValue = modelValue.Elem()

	for i := 0; i < dtoType.NumField(); i++ {
		fieldName := dtoType.Field(i).Name
		dtoFieldValue := dtoValue.FieldByName(fieldName)
		modelFieldValue := modelValue.FieldByName(fieldName)
		if dtoFieldValue.IsValid() && dtoFieldValue.Interface() != reflect.Zero(dtoFieldValue.Type()).Interface() && modelFieldValue.CanSet() {
			modelFieldValue.Set(dtoFieldValue)
		}
	}
}
