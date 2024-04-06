package salesDTO

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
)

type QueryTotalSalesSoldDTO struct {
	appDto.PaginationDTO
	Month  int    `validate:"omitempty,gte=1,lte=12" json:"month" query:"month"`
	Year   int    `validate:"omitempty,gte=1970" json:"year" query:"year"`
	Order  string `validate:"omitempty" json:"order" query:"order"`
	Status string `validate:"omitempty,status" json:"status" query:"status" pg:"operation=equal"`
}
