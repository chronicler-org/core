package managerRouter

import (
	"github.com/chronicler-org/core/src/manager/controller"
	"github.com/chronicler-org/core/src/manager/repository"
	"github.com/chronicler-org/core/src/manager/service"
	"github.com/gofiber/fiber/v2"
)

func NewManagerRouter() *fiber.App {
  router := fiber.New()

  
  repository := managerRepository.InitManagerRepository()
  service := managerService.InitManagerService(repository)
  controller := managerController.InitManagerController(service)

  router.Get("/", controller.HandleFindAll)
  router.Get(":id", controller.HandleFindByID)
  router.Post("/", controller.HandleCreateManager)
  router.Patch(":id", controller.HandleUpdateManager)
  router.Delete(":id", controller.HandleDeleteManager)

  return router
} 
