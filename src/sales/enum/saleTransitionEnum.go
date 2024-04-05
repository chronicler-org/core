package saleEnum

import (
	"github.com/go-playground/validator/v10"
)

type Transition string

const (
	PAGAMENTO_CONFIRMADO Transition = "Pagamento confirmado"
	CONCLUIR_COMPRA      Transition = "Concluir compra"
	CANCELAR_COMPRA      Transition = "Cancelar compra"
)

func validateTransition(fl validator.FieldLevel) bool {
	transition := fl.Field().String()
	validTransition := []string{
		string(PAGAMENTO_CONFIRMADO),
		string(CONCLUIR_COMPRA),
		string(COMPRA_CANCELADA),
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
