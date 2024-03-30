package customerCareDTO

type UpdateCustomerCareeEvaluationDTO struct {
	Score       uint8  `validate:"omitempty,number,min=0,max=5" json:"score"`
	Description string `validate:"omitempty,min=10,max=300" json:"description"`
}
