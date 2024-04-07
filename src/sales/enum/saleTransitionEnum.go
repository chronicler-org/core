package saleEnum

import (
	"github.com/go-playground/validator/v10"
)

type Transition string

const (
	PAYMENT_CONFIRMED Transition = "Pagamento confirmado"
	COMPLETE_PURCHASE      Transition = "Concluir compra"
	CANCEL_PURCHASE      Transition = "Cancelar compra"
)

func validateTransition(fl validator.FieldLevel) bool {
	transition := fl.Field().String()
	validTransition := []string{
		string(PAYMENT_CONFIRMED),
		string(COMPLETE_PURCHASE),
		string(CANCELLED_PURCHASE),
	}

	for _, checkedTransition := range validTransition {
		if checkedTransition == transition {
			return true
		}
	}

	return false
}

func RegisterTransitionValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("transition", validateTransition)
	if err != nil {
		return err
	}
	return err
}
