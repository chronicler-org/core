package managerDTO

type UpdateManagerPasswordDTO struct {
	NewPassword        string `validate:"required,min=8" json:"new_password"`
	ConfirmNewPassword string `validate:"required,min=8" json:"confirm_new_password"`
}
