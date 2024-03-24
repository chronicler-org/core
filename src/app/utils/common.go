package appUtil

import (
	"errors"
	"strings"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appException "github.com/chronicler-org/core/src/app/exceptions"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

func Controller(handler func(*fiber.Ctx) (PaginateResponse, error)) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		result, err := handler(c)

		if err != nil {
			var httpException *appException.HttpException
			var paginateError PaginateErrorResponse

			if errors.As(err, &httpException) {
				paginateError = PaginateError(httpException.GetErrors())
				return c.Status(httpException.GetStatusCode()).JSON(paginateError)

			}

			var pgError *pgconn.PgError
			if errors.As(err, &pgError) {
				customError := []appDto.CustomErrorDTO{
					{
						Code:   "DUPLICATE_ERROR",
						Title:  pgError.Message,
						Detail: pgError.Detail,
					},
				}
				if pgError.Code == "23505" { // PostgreSQL error code for unique violation
					customError[0].Code = "DUPLICATE_ERROR"
					paginateError = PaginateError(customError)
					return c.Status(fiber.StatusBadRequest).JSON(paginateError)
				} else {
					customError[0].Code = "FALLBACK_ERROR"
					paginateError = PaginateError(customError)
					return c.Status(fiber.StatusInternalServerError).JSON(paginateError)
				}
			}

			if errInfo, found := gormErrors[err]; found {
				code := strings.ReplaceAll(strings.ToUpper(errInfo.Message), " ", "_")

				customError := []appDto.CustomErrorDTO{
					{
						Code:   code,
						Title:  errInfo.Message,
						Detail: errInfo.Message,
					},
				}
				return c.Status(errInfo.Status).JSON(PaginateError(customError))
			}

			internalServerException := appException.InternalServerErrorException(err.Error())
			paginateError = PaginateError(internalServerException.GetErrors())
			return c.Status(fiber.StatusInternalServerError).JSON(paginateError)
		}
		return c.Status(200).JSON(result)
	}
}
