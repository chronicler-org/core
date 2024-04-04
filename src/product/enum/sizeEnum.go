package productEnum

import (
	"github.com/go-playground/validator/v10"
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

func RegisterSizeValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("size", validateSize)
	if err != nil {
		return err
	}

	return err
}
