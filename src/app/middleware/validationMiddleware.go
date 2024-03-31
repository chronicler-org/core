package appMiddleware

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
)

func Validate(Validator *validator.Validate) func(interface{}, interface{}) func(*fiber.Ctx) error {
	return func(bodyDto, queryDto interface{}) func(*fiber.Ctx) error {
		return func(c *fiber.Ctx) error {

			if bodyDto != nil {
				c.BodyParser(&bodyDto)
				if err := Validator.StructCtx(c.Context(), bodyDto); err != nil {
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
			}

			if queryDto != nil {
				c.QueryParser(&queryDto)
				if err := Validator.Struct(queryDto); err != nil {
					errors := []appDto.CustomErrorDTO{}
					for _, err := range err.(validator.ValidationErrors) {
						field := err.Field()
						description := err.Tag()

						customError := appDto.CustomErrorDTO{
							Code:   "INVALID_QUERY",
							Title:  fmt.Sprintf("Campo %s na query é inválido", field),
							Detail: description,
						}
						errors = append(errors, customError)
					}

					return c.Status(http.StatusBadRequest).JSON(appUtil.PaginateError(errors))
				}
			}

			return c.Next()
		}

	}
}
