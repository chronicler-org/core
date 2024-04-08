package appUtil

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateBirthDateFormat(fl validator.FieldLevel) bool {
	birthDateString := fl.Field().String()
	_, err := time.Parse(time.RFC3339, birthDateString)
	return err == nil
}

func RegisterBirthDateFormatValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("birthdate", ValidateBirthDateFormat)
	if err != nil {
		return err
	}

	return err
}

