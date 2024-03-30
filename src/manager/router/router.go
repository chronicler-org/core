package managerRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
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
) {
	router.Get("/manager", middleware.Validate(nil, &managerDTO.QueryManagerDTO{}), appUtil.Controller(managerController.HandleFindAll))
	router.Get("/manager/:id", appUtil.Controller(managerController.HandleFindByID))
	router.Post("/manager", middleware.Validate(&managerDTO.CreateManagerDTO{}, nil), appUtil.Controller(managerController.HandleCreateManager))
	router.Patch("/manager/:id", middleware.Validate(&managerDTO.UpdateManagerDTO{}, nil), appUtil.Controller(managerController.HandleUpdateManager))
	router.Delete("/manager/:id", appUtil.Controller(managerController.HandleDeleteManager))
}
