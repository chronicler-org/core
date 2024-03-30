package authRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
	authController "github.com/chronicler-org/core/src/auth/controller"
	authDTO "github.com/chronicler-org/core/src/auth/dto"
	authService "github.com/chronicler-org/core/src/auth/service"
	managerService "github.com/chronicler-org/core/src/manager/service"
)

func InitAuthModule(
	db *gorm.DB,
	managerServ *managerService.ManagerService,
	attendantServ *attendantService.AttendantService,
) (*authController.AuthController, *authService.AuthService) {
	authServ := authService.InitAuthService(managerServ, attendantServ)
	authCtrl := authController.InitAuthrController(authServ)

	return authCtrl, authServ
}

func InitAuthRouter(
	router *fiber.App,
	authController *authController.AuthController,
) {
	router.Post("/auth/login", middleware.Validate(&authDTO.AuthLoginDTO{}, nil), appUtil.Controller(authController.HandleLogin))
}
