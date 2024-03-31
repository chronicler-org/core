package appMiddleware

import (
	"github.com/gofiber/fiber/v2"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	authExceptionMessage "github.com/chronicler-org/core/src/auth/messages"
)

func RouteAccessMiddleware(allowedRoles []authEnum.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userRole authEnum.Role

		if c.Locals(authEnum.ManagerRole) != nil {
			userRole = authEnum.ManagerRole
		} else {
			userRole = authEnum.AttendantRole
		}

		authorized := false
		for _, role := range allowedRoles {
			if userRole == role {
				authorized = true
				break
			}
		}

		if !authorized {
			exception := appException.ForbiddenException(authExceptionMessage.USER_SERVICE_ACCESS_INFO_DENIED)
			return c.Status(exception.GetStatusCode()).JSON(appUtil.PaginateError(exception.GetErrors()))
		}

		return c.Next()
	}
}
