package authMiddleware

import (
	"os"
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	authExceptionMessage "github.com/chronicler-org/core/src/auth/messages"
	managerService "github.com/chronicler-org/core/src/manager/service"
)

func WithAuth(
	managerService *managerService.ManagerService,
	attendantService *attendantService.AttendantService,
) fiber.Handler {
	authSecret := []byte(os.Getenv("AT_SECRET"))

	exceptionInvalidAt := appException.UnauthorizedException(authExceptionMessage.INVALID_AT)
	exceptionExpiredAt := appException.UnauthorizedException(authExceptionMessage.EXPIRED_AT)

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: authSecret,
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			tokenString := strings.Replace(c.Get("Authorization"), "Bearer ", "", 1)
			token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				return nil, nil
			})
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return c.Status(exceptionInvalidAt.GetStatusCode()).JSON(appUtil.PaginateError(exceptionInvalidAt.GetErrors()))
			}
			id := claims["sub"].(string)
			role := claims["role"].(string)

			if role == authEnum.ManagerRole {
				manager, err := managerService.FindByID(id)
				if err != nil {
					return c.Status(exceptionExpiredAt.GetStatusCode()).JSON(appUtil.PaginateError(exceptionExpiredAt.GetErrors()))
				}
				c.Locals("manager", manager)
			} else {
				attendant, err := attendantService.FindAttendantByID(id)
				if err != nil {
					return c.Status(exceptionExpiredAt.GetStatusCode()).JSON(appUtil.PaginateError(exceptionExpiredAt.GetErrors()))
				}
				c.Locals("attendant", attendant)
			}
			return c.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(exceptionInvalidAt.GetStatusCode()).JSON(appUtil.PaginateError(exceptionInvalidAt.GetErrors()))
			}
			return c.Status(exceptionExpiredAt.GetStatusCode()).JSON(appUtil.PaginateError(exceptionExpiredAt.GetErrors()))
		},
		TokenLookup: "header:Authorization",
		AuthScheme:  "Bearer",
	})
}
