package authRouter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
	authController "github.com/chronicler-org/core/src/auth/controller"
	authDTO "github.com/chronicler-org/core/src/auth/dto"
	authService "github.com/chronicler-org/core/src/auth/service"
	managerService "github.com/chronicler-org/core/src/manager/service"
)

func InitAuthRouter(
	router *fiber.App,
	managerServ *managerService.ManagerService,
	attendantServ *attendantService.AttendantService,
) {
	authService := authService.InitAuthService(managerServ, attendantServ)
	authController := authController.InitAuthrController(authService)

	router.Post("/auth/login", middleware.Validate(&authDTO.AuthLoginDTO{}, nil), appUtil.Controller(authController.HandleLogin))
}
