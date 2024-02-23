package managerRouter

import (
	managerController "github.com/chronicler-org/core/src/manager/controller"
	"github.com/gofiber/fiber/v2"
)

func NewManagerRouter() *fiber.App {
  router := fiber.New()

  router.Get("/", managerController.HandleGetAll)
  router.Get(":id", managerController.HandleGetById)
  router.Post("/", managerController.HandleCreateManager)
  router.Patch(":id", managerController.HandleUpdateManager)
  router.Delete(":id", managerController.HandleDeleteManager)

  return router
} 
