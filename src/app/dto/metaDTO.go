package dto

import (
	"encoding/json"
	"time"
)

type MetaDTO struct {
	Count           int       `json:"count"`
	Page            int       `json:"page"`
	TotalPages      int       `json:"totalPages"`
	TotalCount      int       `json:"totalCount"`
	Limit           int       `json:"limit"`
	RequestDateTime time.Time `json:"requestDateTime"`
}