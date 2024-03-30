package authDTO

type AuthLoginDTO struct {
	Email    string `validate:"required,email,max=50" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
