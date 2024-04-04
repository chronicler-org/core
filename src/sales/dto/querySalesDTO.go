package salesDTO

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QuerySalesDTO struct {
	appDto.PaginationDTO
	CustomerCareID string `validate:"omitempty,uuid" query:"customer_care_id"`
}
