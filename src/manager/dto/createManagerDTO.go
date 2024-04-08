package managerDTO

type CreateManagerDTO struct {
	CPF       string `validate:"required,cpf" json:"cpf"`
	Name      string `validate:"required,min=10,max=50" json:"name"`
	Email     string `validate:"required,email,max=50" json:"email"`
	Password  string `validate:"required,min=8" json:"password"`
	TeamId    string `validate:"omitempty,uuid" json:"team_id"`
	BirthDate string `validate:"required,birthdate,notfuturedate" json:"birth_date"`
}
