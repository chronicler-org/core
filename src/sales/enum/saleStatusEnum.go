package saleEnum

import "github.com/go-playground/validator/v10"

type SaleStatus string

const (
	AWAITING_PAYMENT   SaleStatus = "Aguardando pagamento"
	PURCHASE_CONFIRMED SaleStatus = "Compra confirmada"
	PURCHASE_COMPLETED SaleStatus = "Compra concluida"
	CANCELLED_PURCHASE SaleStatus = "Compra cancelada"
)

func validateStatus(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	validStatus := []string{
		string(AWAITING_PAYMENT),
		string(PURCHASE_CONFIRMED),
		string(PURCHASE_COMPLETED),
		string(CANCELLED_PURCHASE),
	}

	for _, checkedStatus := range validStatus {
		if checkedStatus == status {
			return true
		}
	}

	return false
}

func RegisterStatusValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("status", validateStatus)
	if err != nil {
		return err
	}
	return err
}
