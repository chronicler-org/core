package appUtil

import (
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
)

func ValidateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	return cpfcnpj.ValidateCPF(cpf)
}

func RegisterCPFValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("cpf", ValidateCPF)
	if err != nil {
		return err
	}

	return err
}
