package customerDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type CustomerQueryDTO struct {
	appDto.PaginationDTO
	CreatedMonth int `validate:"omitempty,min=1,max=12" json:"month" query:"month" `
	CreatedYear  int `validate:"omitempty,number" json:"year" query:"year" `
}
