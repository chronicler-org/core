package util

import (
	"encoding/json"
	"time"
	"math"
)

type MetaData struct {
	Count           int       `json:"count"`
	Page            int       `json:"page"`
	TotalPages      int       `json:"totalPages"`
	TotalCount      int       `json:"totalCount"`
	Limit           int       `json:"limit"`
	RequestDateTime time.Time `json:"requestDateTime"`
}

type ApiResponse struct {
	Meta            MetaData       `json:"meta"`
	Result          []interface{}  `json:"result"`
}

func Paginate(data []interface{}, page int, totalCount int, limit int) ApiResponse {
	meta := MetaData {
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

func PaginateSingle(data interface{}) ApiResponse {
	meta := MetaData {
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

