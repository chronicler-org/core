package appUtil

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func ValidateNotFutureDate(fl validator.FieldLevel) bool {
	dateString := fl.Field().String()
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		return false
	}
	now := time.Now()
	return date.Before(now)
}

func RegisterNotFutureDateValidation(validate *validator.Validate) error {
	err := validate.RegisterValidation("notfuturedate", ValidateNotFutureDate)
	if err != nil {
		return err
	}

	return err
}