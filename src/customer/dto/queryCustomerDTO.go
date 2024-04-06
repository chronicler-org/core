package customerDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerDTO struct {
	appDto.PaginationDTO
	CreatedMonth int    `validate:"omitempty,min=1,max=12" query:"created_month" pg:"operation=equal,date_extract=MONTH,name=created_at" `
	CreatedYear  int    `validate:"omitempty,number" query:"created_year" pg:"operation=equal,date_extract=YEAR,name=created_at" `
	CPF          string `validate:"omitempty,cpf" query:"cpf" pg:"operation=like" `
	Order        string `validate:"omitempty" json:"order" query:"order"`
}
