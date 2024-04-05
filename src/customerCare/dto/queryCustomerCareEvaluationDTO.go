package customerCareDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerCareEvaluationDTO struct {
	appDto.PaginationDTO
	CustomerCareID string `validate:"omitempty,uuid" query:"customer_care_id" pg:"operation=equal"`
	CustomerCPF    string `validate:"omitempty,cpf" query:"customer_cpf" pg:"operation=like"`
}
