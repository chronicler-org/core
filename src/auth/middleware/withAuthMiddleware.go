package authMiddleware

import (
	"fmt"
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authExceptionMessage "github.com/chronicler-org/core/src/auth/messages"
)

func WithAuth() fiber.Handler {
	authSecret := []byte(os.Getenv("AT_SECRET"))

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: authSecret,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			token := c.Get("Authorization")
			fmt.Println(token)
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				exception := appException.UnauthorizedException(authExceptionMessage.INVALID_AT)
				return c.Status(exception.GetStatusCode()).JSON(appUtil.PaginateError(exception.GetErrors()))
			}
			exception := appException.UnauthorizedException(authExceptionMessage.EXPIRED_AT)
			return c.Status(exception.GetStatusCode()).JSON(appUtil.PaginateError(exception.GetErrors()))
		},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	})
}
