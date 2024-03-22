package appUtil

import (
	"errors"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	"github.com/gofiber/fiber/v2"
)

func Controller(handler func(*fiber.Ctx) (PaginateResponse, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		result, err := handler(c)

		if err != nil {
			target := &appException.HttpException{}

			if errors.As(err, &target) {
				paginateError := PaginateError(target.GetErrors())
				c.Status(target.GetStatusCode()).JSON(paginateError)
			} else {
				internalServerException := appException.InternalServerErrorException(err.Error())
				paginateError := PaginateError(internalServerException.GetErrors())
				c.Status(fiber.StatusInternalServerError).JSON(paginateError)
			}
		}
		return c.Status(200).JSON(result)
	}
}