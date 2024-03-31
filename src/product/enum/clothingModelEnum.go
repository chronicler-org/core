package productEnum

import (
	appUtil "github.com/chronicler-org/core/src/app/utils"
	ut "github.com/go-playground/universal-translator"
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

func RegisterModelValidationAndTranslation(validate *validator.Validate, trans ut.Translator) error {
	err := validate.RegisterValidation("model", validateModel)
	if err != nil {
		return err
	}

	validModels := []interface{}{
		string(TShirt), string(Shirt), string(Pants), string(Skirt), string(Dress),
		string(Jacket), string(Sportswear), string(Underwear), string(Swimwear), string(FormalWear),
	}

	err = validate.RegisterTranslation("model", trans, func(ut ut.Translator) error {
		return ut.Add("model", appUtil.GenerateEnumErrorDetail("Model", validModels), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("model", fe.Field())
		return t
	})

	return err
}
