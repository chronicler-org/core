package managerRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	managerController "github.com/chronicler-org/core/src/manager/controller"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerRepository "github.com/chronicler-org/core/src/manager/repository"
	managerService "github.com/chronicler-org/core/src/manager/service"
	teamService "github.com/chronicler-org/core/src/team/service"
)

func InitManagerModule(
	db *gorm.DB,
	teamServ *teamService.TeamService,
) (*managerController.ManagerController, *managerService.ManagerService) {
	managerRepo := managerRepository.InitManagerRepository(db)
	managerServ := managerService.InitManagerService(managerRepo, teamServ)
	managerCtrl := managerController.InitManagerController(managerServ)

	return managerCtrl, managerServ
}

func InitManagerRouter(
	router *fiber.App,
	managerController *managerController.ManagerController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	managerRouter := router.Group("/manager")
	managerRouter.Use(appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole}))

	managerRouter.Get("/",
		validatorMiddleware(nil, &managerDTO.QueryManagerDTO{}),
		appUtil.Controller(managerController.HandleFindAll),
	)
	managerRouter.Get("/me",
		appUtil.Controller(managerController.HandleGetLoggedManager),
	)
	managerRouter.Get("/:id",
		appUtil.Controller(managerController.HandleFindByID),
	)
	managerRouter.Post("/",
		validatorMiddleware(&managerDTO.CreateManagerDTO{}, nil),
		appUtil.Controller(managerController.HandleCreateManager),
	)
	managerRouter.Patch("/updatePassword",
		validatorMiddleware(&managerDTO.UpdateManagerPasswordDTO{}, nil),
		appUtil.Controller(managerController.HandleUpdateManagerPassword),
	)
	managerRouter.Patch("/:id",
		validatorMiddleware(&managerDTO.UpdateManagerDTO{}, nil),
		appUtil.Controller(managerController.HandleUpdateManager),
	)
	managerRouter.Delete("/:id",
		appUtil.Controller(managerController.HandleDeleteManager),
	)
}
