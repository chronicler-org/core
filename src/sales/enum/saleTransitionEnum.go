package saleEnum

import (
	appUtil "github.com/chronicler-org/core/src/app/utils"
	ut "github.com/go-playground/universal-translator"
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

func RegisterTransitionValidationAndTranslation(validate *validator.Validate, translator ut.Translator) error {
	err := validate.RegisterValidation("transition", validateTransition)
	if err != nil {
		return err
	}

	validTransition := []interface{}{
		string(PAGAMENTO_CONFIRMADO),
		string(CONCLUIR_COMPRA),
		string(CANCELAR_COMPRA),
	}

	err = validate.RegisterTranslation("transition", translator, func(ut ut.Translator) error {
		return ut.Add("transition", appUtil.GenerateEnumErrorDetail("Transition", validTransition), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("transition", fe.Field())
		return t
	})

	return err
}
