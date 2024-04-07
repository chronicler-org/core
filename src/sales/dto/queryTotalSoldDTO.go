package salesDTO

import (
	"time"

	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QueryTotalSalesSoldDTO struct {
	appDto.PaginationDTO
	StartDate time.Time `validate:"omitempty" json:"start_date" query:"start_date"`
	EndDate   time.Time `validate:"omitempty" json:"end_date" query:"end_date"`
}
