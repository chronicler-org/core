package saleEnum

import "github.com/go-playground/validator/v10"

type PaymentMethod string

const (
	CREDIT_CARD PaymentMethod = "Cartão de Crédito"
	DEBIT_CARD  PaymentMethod = "Cartão de Débito"
	CASH        PaymentMethod = "Dinheiro"
)

func validatePaymentMethod(fl validator.FieldLevel) bool {
	paymentMethod := fl.Field().String()
	validPaymentMethod := []string{
		string(CREDIT_CARD),
		string(DEBIT_CARD),
		string(CASH),
	}

	for _, checkedPaymentMethod := range validPaymentMethod {
		if checkedPaymentMethod == paymentMethod {
			return true
		}
	}

	return false
}

func RegisterPaymentMethodValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("paymentMethod", validatePaymentMethod)
	if err != nil {
		return err
	}
	return err
}
