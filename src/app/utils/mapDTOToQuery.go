package appUtil

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

func MapDTOToQuery(dto interface{}, query *gorm.DB) (*gorm.DB, appDto.PaginationDTO) {
	var paginationDTO appDto.PaginationDTO

	dtoValue := reflect.ValueOf(dto)
	if dtoValue.Kind() == reflect.Ptr {
		dtoValue = dtoValue.Elem()
	}
	dtoType := dtoValue.Type()

	for i := 0; i < dtoType.NumField(); i++ {

		fieldName := dtoType.Field(i).Name
		fieldValue := dtoValue.Field(i)

		switch fieldName {
		case "PaginationDTO":
			paginationDTOValue := fieldValue.Interface().(appDto.PaginationDTO)
			paginationDTO.Limit = paginationDTOValue.Limit
			paginationDTO.Page = paginationDTOValue.Page
		case "Limit":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int {
				paginationDTO.Limit = int(fieldValue.Int())
			}
		case "Page":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int {
				paginationDTO.Page = int(fieldValue.Int())
			}
		case "CreatedMonth":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int && fieldValue.Int() != 0 {
				query = query.Where("EXTRACT(MONTH FROM created_at) = ?", fieldValue.Int())
			}
		case "CreatedYear":
			if fieldValue.IsValid() && fieldValue.Type().Kind() == reflect.Int && fieldValue.Int() != 0 {
				query = query.Where("EXTRACT(YEAR FROM created_at) = ?", int(fieldValue.Int()))
			}
		default:

			fmt.Println(fieldValue.Interface())
			if fieldValue.Interface() != "" && !fieldValue.IsZero() {

				tag := dtoType.Field(i).Tag.Get("query")
				if tag != "" {
					query = query.Where(fmt.Sprintf("%s = ?", tag), fieldValue.Interface())
				}
			}
		}

	}

	return query, paginationDTO
}
