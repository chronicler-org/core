package dto

import (
	"encoding/json"
	"time"
	appUtils "github.com/chronicler-org/core/src/app/utils"
)

type CustomError struct {
	Code            string    `json:"code"`
	Title           string    `json:"title"`
	Detail          int       `json:"detail"`
}

type ApiErrorResponse struct {
	Meta            MetaData       `json:"meta"`
	Errors          []CustomError  `json:"errors"`
}

func ResponseError(errors []CustomError) ApiErrorResponse {
	meta := appUtils.MetaData {
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