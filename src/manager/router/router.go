package managerRouter

import (
	appUtil "github.com/chronicler-org/core/src/app/utils"
	"github.com/chronicler-org/core/src/manager/controller"
	"github.com/chronicler-org/core/src/manager/repository"
	"github.com/chronicler-org/core/src/manager/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitManagerRouter(router *fiber.App, db *gorm.DB) {

	repository := managerRepository.InitManagerRepository(db)
	validate := validator.New()
	service := managerService.InitManagerService(repository, validate)
	controller := managerController.InitManagerController(service)

	router.Get("/manager", controller.HandleFindAll)
	router.Get("/manager/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/manager", controller.HandleCreateManager)
	router.Patch("/manager/:id", controller.HandleUpdateManager)
	router.Delete("/manager/:id", controller.HandleDeleteManager)

}
