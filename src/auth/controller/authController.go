package authController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	authDTO "github.com/chronicler-org/core/src/auth/dto"
	authService "github.com/chronicler-org/core/src/auth/service"
)

type AuthController struct {
	authService *authService.AuthService
}

func InitAuthrController(authService *authService.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (controller *AuthController) HandleLogin(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var authLoginDTO authDTO.AuthLoginDTO
	c.BodyParser(&authLoginDTO)

	result, err := controller.authService.Login(authLoginDTO)
	return appUtil.PaginateSingle(result), err
}
