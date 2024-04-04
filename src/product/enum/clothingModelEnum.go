package productEnum

import (
	"github.com/go-playground/validator/v10"
)

type ClothingModel string

const (
	TShirt     ClothingModel = "Camiseta"
	Shirt      ClothingModel = "Camisa"
	Pants      ClothingModel = "Calça"
	Skirt      ClothingModel = "Saia"
	Dress      ClothingModel = "Vestido"
	Jacket     ClothingModel = "Casaco"
	Sportswear ClothingModel = "Roupas Esportivas"
	Underwear  ClothingModel = "Roupa Íntima"
	Swimwear   ClothingModel = "Roupa de Banho"
	FormalWear ClothingModel = "Roupa Formal"
)

func validateModel(fl validator.FieldLevel) bool {
	model := fl.Field().String()
	validModels := []string{
		string(TShirt), string(Shirt), string(Pants), string(Skirt), string(Dress),
		string(Jacket), string(Sportswear), string(Underwear), string(Swimwear), string(FormalWear),
	}

	for _, validModel := range validModels {
		if model == validModel {
			return true
		}
	}

	return false
}

func RegisterModelValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("model", validateModel)
	if err != nil {
		return err
	}

	return err
}
