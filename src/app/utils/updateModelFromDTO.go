package appUtil

import (
	"reflect"
)

func UpdateModelFromDTO(model interface{}, dto interface{}) {
	modelValue := reflect.ValueOf(model)
	dtoValue := reflect.ValueOf(dto)

	if modelValue.Kind() == reflect.Ptr {
		modelValue = modelValue.Elem()
	} else {
		panic("model must be a pointer to a struct")
	}

	if dtoValue.Kind() == reflect.Ptr {
		dtoValue = dtoValue.Elem()
	} else {
		panic("dto must be a pointer to a struct")
	}

	for i := 0; i < dtoValue.NumField(); i++ {
		dtoField := dtoValue.Type().Field(i)
		dtoFieldValue := dtoValue.Field(i)

		modelField := modelValue.FieldByName(dtoField.Name)
		if !modelField.IsValid() || !modelField.CanSet() {
			continue
		}

		if dtoFieldValue.IsZero() {
			continue
		}

		switch dtoFieldValue.Kind() {
		case reflect.Slice:
			// Check if the field in the model is a slice of strings
			if modelField.Type().Elem().Kind() == reflect.String {
				modelField.Set(dtoFieldValue)
			}
		default:
			modelField.Set(dtoFieldValue)
		}
	}
}
