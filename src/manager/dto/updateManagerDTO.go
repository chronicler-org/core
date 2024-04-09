package managerDTO

type UpdateManagerDTO struct {
	CPF       string `validate:"omitempty,cpf" json:"cpf,omitempty"`
	Name      string `validate:"omitempty,min=10,max=50" json:"name,omitempty"`
	Email     string `validate:"omitempty,email,max=50" json:"email,omitempty"`
	Password  string `validate:"omitempty,min=8" json:"Password,omitempty"`
	TeamId    string `validate:"omitempty,uuid" json:"team_id"`
	BirthDate string `json:"birth_date,omitempty"`
}
