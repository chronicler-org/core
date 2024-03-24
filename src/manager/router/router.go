package managerRouter

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	managerController "github.com/chronicler-org/core/src/manager/controller"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerRepository "github.com/chronicler-org/core/src/manager/repository"
	managerService "github.com/chronicler-org/core/src/manager/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitManagerRouter(router *fiber.App, db *gorm.DB) {

	repository := managerRepository.InitManagerRepository(db)
	validate := validator.New()
	service := managerService.InitManagerService(repository, validate)
	controller := managerController.InitManagerController(service)

	router.Get("/manager", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(controller.HandleFindAll))
	router.Get("/manager/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/manager", middleware.Validate(&managerDTO.CreateManagerDTO{}, nil), appUtil.Controller(controller.HandleCreateManager))
	router.Patch("/manager/:id", controller.HandleUpdateManager)
	router.Delete("/manager/:id", controller.HandleDeleteManager)

}
