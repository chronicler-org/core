package productEnum

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"

	appUtil "github.com/chronicler-org/core/src/app/utils"
)

type Size string

const (
	SizePP Size = "PP"
	SizeP  Size = "P"
	SizeM  Size = "M"
	SizeG  Size = "G"
	SizeGG Size = "GG"
)

func validateSize(fl validator.FieldLevel) bool {
	size := fl.Field().String()
	validSizes := []string{
		string(SizePP),
		string(SizeP),
		string(SizeM),
		string(SizeG),
		string(SizeGG),
	}
	for _, validSize := range validSizes {
		if size == validSize {
			return true
		}
	}

	return false
}

func RegisterSizeValidationAndTranslation(validate *validator.Validate, trans ut.Translator) error {
	err := validate.RegisterValidation("size", validateSize)
	if err != nil {
		return err
	}
	validSizes := []interface{}{
		string(SizePP), string(SizeP), string(SizeM), string(SizeG), string(SizeG),
	}

	err = validate.RegisterTranslation("size", trans, func(ut ut.Translator) error {
		return ut.Add("size", appUtil.GenerateEnumErrorDetail("Size", validSizes), true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("size", fe.Field())
		return t
	})

	return err
}
