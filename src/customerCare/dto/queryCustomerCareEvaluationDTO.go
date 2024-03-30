package customerCareDTO

import appDto "github.com/chronicler-org/core/src/app/dto"

type QueryCustomerCareEvaluationDTO struct {
	appDto.PaginationDTO
	CustomerCareID string `validate:"omitempty,uuid" query:"customer_care_id"`
	CustomerCPF    string `validate:"omitempty,number,len=11" query:"customer_cpf"`
}
