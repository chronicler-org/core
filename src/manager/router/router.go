package managerRouter

import (
	"github.com/chronicler-org/core/src/manager/controller"
	"github.com/chronicler-org/core/src/manager/repository"
	"github.com/chronicler-org/core/src/manager/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func NewManagerRouter() *fiber.App {
  router := fiber.New()

  
  repository := managerRepository.InitManagerRepository()
  validate := validator.New()
  service := managerService.InitManagerService(repository, validate)
  controller := managerController.InitManagerController(service)

  router.Get("/", controller.HandleFindAll)
  router.Get(":id", controller.HandleFindByID)
  router.Post("/", controller.HandleCreateManager)
  router.Patch(":id", controller.HandleUpdateManager)
  router.Delete(":id", controller.HandleDeleteManager)

  return router
} 
