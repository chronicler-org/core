package tagRouter

import (
	tagController "github.com/chronicler-org/core/src/tag/controller"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	tagService "github.com/chronicler-org/core/src/tag/service"
	"github.com/gofiber/fiber/v2"
)

func NewTagRouter() *fiber.App {
	router := fiber.New()

	repository := tagRepository.InitTagRepository()
	service := tagService.InitTagService(repository)
	controller := tagController.InitTagController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get(":id", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateTag)
	router.Patch("/:id", controller.HandleUpdateTag)
	router.Delete("/:id", controller.HandleDeleteTag)

	return router
}
