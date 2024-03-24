package middleware

import (
	"fmt"
	"net/http"

	"github.com/go-playground/locales/pt_BR"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	ut "github.com/go-playground/universal-translator"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
)

func Validate(dto interface{}) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		Validator := validator.New()

		pt := pt_BR.New()
		uni := ut.New(pt, pt)
		trans, _ := uni.GetTranslator("pt_BR")

		// Registro da validação de CPF
		appUtil.RegisterCPFValidationAndTranslation(Validator, trans)

		c.BodyParser(&dto)

		if err := Validator.StructCtx(c.Context(), dto); err != nil {
			errors := []appDto.CustomErrorDTO{}
			for _, err := range err.(validator.ValidationErrors) {
				field := err.Field()
				description := err.Tag()

				customError := appDto.CustomErrorDTO{
					Code:   "INVALID_DATA",
					Title:  fmt.Sprintf("Campo %s é inválido", field),
					Detail: description,
				}
				errors = append(errors, customError)
			}

			return c.Status(http.StatusBadRequest).JSON(appUtil.PaginateError(errors))
		}

		return c.Next()
	}

}
