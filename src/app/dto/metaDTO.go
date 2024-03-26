package appDto

type MetaDTO struct {
	Count      int `json:"count"`
	Page       int `json:"page"`
	TotalPages int `json:"total_pages"`
	TotalCount int64 `json:"total_count"`
	Limit      int `json:"limit"`
}
