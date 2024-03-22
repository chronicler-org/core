package util

import (
	"encoding/json"
	"time"
	"math"
	"github.com/chronicler-org/core/src/app/dto"
)

type PaginateResponse struct {
	Meta            appDTO.MetaDTO `json:"meta"`
	Result          []interface{}  `json:"result"`
}

func Paginate(data []interface{}, page int, totalCount int, limit int) PaginateResponse {
	meta := MetaDTO {
		Count: len(data),
		Page: page,
		TotalPages: int(math.Ceil(float64(totalCount) / float64(limit))),
		TotalCount: totalCount,
		Limit: limit,
		RequestDateTime: time.Now(),
	}

	return ApiResponse {
		Meta: meta,
		Result: data
	}
}

func PaginateSingle(data interface{}) PaginateResponse {
	meta := MetaDTO {
		Count: 1,
		Page: 1,
		TotalPages: 1,
		TotalCount: 1,
		Limit: 1,
		RequestDateTime: time.Now(),
	}

	return ApiResponse {
		Meta: meta,
		Result: data
	}
}

type PaginateErrorResponse struct {
	Meta            MetaDTO           `json:"meta"`
	Errors          []CustomErrorDTO  `json:"errors"`
}

func PaginateError(errors []CustomErrorDTO) PaginateErrorResponse {
	meta := MetaDTO {
		Count: len(errors),
		Page: 1,
		TotalPages: 1,
		TotalCount: len(errors),
		Limit: 1,
		RequestDateTime: time.Now(),
	}

	return ApiErrorResponse {
		Meta: meta,
		errors: errors
	}
}
