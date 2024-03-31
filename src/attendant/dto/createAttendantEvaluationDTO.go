package attendantDTO

type CreateAttendantEvaluationDTO struct {
	Score       uint8  `validate:"required,number,min=0,max=5" json:"score"`
	Description string `validate:"required,min=10,max=300" json:"description"`
	AvaluatedID string `validate:"omitempty,uuid" json:"avaluated_id"`
}
