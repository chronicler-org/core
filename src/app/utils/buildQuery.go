package appUtil

import (
	"fmt"
	"reflect"
	"strings"

	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type TQueryBuilder struct {
	dto           interface{}
	query         *gorm.DB
	order         string
	paginationDTO appDto.PaginationDTO
}

func QueryBuilder(dto interface{}, query *gorm.DB) *TQueryBuilder {
	return &TQueryBuilder{
		dto:           dto,
		query:         query,
		paginationDTO: appDto.PaginationDTO{},
	}
}

func (qb *TQueryBuilder) BuildQuery() *gorm.DB {
	dtoValue := reflect.ValueOf(qb.dto)
	if dtoValue.Kind() == reflect.Ptr {
		dtoValue = dtoValue.Elem()
	}
	dtoType := dtoValue.Type()

	for i := 0; i < dtoType.NumField(); i++ {
		field := dtoType.Field(i)
		value := dtoValue.Field(i)
		tags := parseTags(field.Tag.Get("pg"))

		// Construct field name
		queryString := ""
		var fieldName string
		if tags["name"] != "" {
			fieldName = tags["name"]
		} else if field.Tag.Get("query") != "" {
			fieldName = field.Tag.Get("query")
		} else {
			fieldName = field.Name
		}

		switch fieldName {
		case "PaginationDTO":
			paginationDTOValue := value.Interface().(appDto.PaginationDTO)
			qb.paginationDTO.Limit = paginationDTOValue.Limit
			qb.paginationDTO.Page = paginationDTOValue.Page
		case "Limit":
			if value.IsValid() && value.Type().Kind() == reflect.Int {
				qb.paginationDTO.Limit = int(value.Int())
			}
		case "Page":
			if value.IsValid() && value.Type().Kind() == reflect.Int {
				qb.paginationDTO.Page = int(value.Int())
			}
		case "order":
			if value.IsValid() && value.Type().Kind() == reflect.String && value.String() != "" {
				qb.order = value.String()
			}
		default:

			if tags["date_extract"] != "" && value.Int() != 0 {
				queryString = buildQueryString(fmt.Sprintf("EXTRACT(%s FROM %s)", tags["date_extract"], fieldName), tags)
				qb.query = qb.query.Where(queryString, int(value.Int()))
			} else if tags["date_extract"] == "" && value.Interface() != "" {
				queryString = buildQueryString(fieldName, tags)

				if tags["operation"] == "like" && value.Type().Kind() == reflect.String {
					qb.query = qb.query.Where(queryString, fmt.Sprintf("%%%s%%", value.Interface()))
				} else {
					qb.query = qb.query.Where(queryString, value.Interface())
				}
			}
		}

	}

	return qb.query
}

func (qb *TQueryBuilder) GetPagination() (int, int) {
	page := qb.paginationDTO.GetPage()
	limit := qb.paginationDTO.GetLimit()
	offset := (page - 1) * limit

	return offset, limit
}

func (qb *TQueryBuilder) ApplyOrder() *gorm.DB {
	for fieldName, direction := range qb.getOrder() {
		qb.query = qb.query.Order(fmt.Sprintf("%s %s", fieldName, direction))
	}
	return qb.query
}

func (qb *TQueryBuilder) getOrder() map[string]string {
	orderMap := make(map[string]string)

	if qb.order != "" {
		orders := strings.Split(qb.order, ",")
		for _, order := range orders {
			orderParts := strings.Split(order, ":")
			if len(orderParts) != 2 {
				continue
			}
			fieldName := strings.TrimSpace(orderParts[0])
			direction := strings.TrimSpace(orderParts[1])

			orderMap[fieldName] = direction
		}
	}

	return orderMap
}

func parseTags(tags string) map[string]string {
	tagMap := make(map[string]string)
	parts := strings.Split(tags, ",")
	for _, part := range parts {
		kv := strings.Split(part, "=")
		if len(kv) != 2 {
			continue
		}
		tagMap[kv[0]] = kv[1]
	}

	return tagMap
}

func buildQueryString(fieldName string, tags map[string]string) string {
	var queryString string

	switch tags["operation"] {
	case "like":
		queryString = fmt.Sprintf("%s LIKE ?", fieldName)
	case "equal":
		queryString = fmt.Sprintf("%s = ?", fieldName)
	case "less":
		queryString = fmt.Sprintf("%s < ?", fieldName)
	case "greater":
		queryString = fmt.Sprintf("%s > ?", fieldName)
	default:
		queryString = fmt.Sprintf("%s = ?", fieldName)
	}

	return queryString
}
