package appUtil

import (
	"math"
	"time"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type PaginateResponse struct {
	Meta   appDto.MetaDTO `json:"meta"`
	Result interface{}  `json:"result"`
}

func Paginate(data []interface{}, page int, totalCount int, limit int) PaginateResponse {
	meta := appDto.MetaDTO{
		Count:           len(data),
		Page:            page,
		TotalPages:      int(math.Ceil(float64(totalCount) / float64(limit))),
		TotalCount:      totalCount,
		Limit:           limit,
		RequestDateTime: time.Now(),
	}

	return PaginateResponse{
		Meta:   meta,
		Result: data,
	}
}

func PaginateSingle(data interface{}) PaginateResponse {
	meta := appDto.MetaDTO{
		Count:           1,
		Page:            1,
		TotalPages:      1,
		TotalCount:      1,
		Limit:           1,
		RequestDateTime: time.Now(),
	}

	return PaginateResponse{
		Meta:   meta,
		Result: data,
	}
}

type PaginateErrorResponse struct {
	Meta   appDto.MetaDTO          `json:"meta"`
	Errors []appDto.CustomErrorDTO `json:"errors"`
}

func PaginateError(errors []appDto.CustomErrorDTO) PaginateErrorResponse {
	meta := appDto.MetaDTO{
		Count:           len(errors),
		Page:            1,
		TotalPages:      1,
		TotalCount:      len(errors),
		Limit:           1,
		RequestDateTime: time.Now(),
	}

	return PaginateErrorResponse{
		Meta:   meta,
		Errors: errors,
	}
}
