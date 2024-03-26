package appUtil

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/klassmann/cpfcnpj"
)

func ValidateCPF(fl validator.FieldLevel) bool {
	cpf := fl.Field().String()
	return cpfcnpj.ValidateCPF(cpf)
}

func RegisterCPFValidationAndTranslation(validate *validator.Validate, trans ut.Translator) error {
	err := validate.RegisterValidation("cpf", ValidateCPF)
	if err != nil {
		return err
	}

	err = validate.RegisterTranslation("cpf", trans, func(ut ut.Translator) error {
		return ut.Add("cpf", "{0} deve ser um CPF válido", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("cpf", fe.Field())
		return t
	})

	return err
}
