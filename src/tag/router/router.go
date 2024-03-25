package tagRouter

import (
	tagController "github.com/chronicler-org/core/src/tag/controller"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	tagService "github.com/chronicler-org/core/src/tag/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitTagRouter(router *fiber.App, db *gorm.DB) {
	repository := tagRepository.InitTagRepository(db)
	service := tagService.InitTagService(repository)
	controller := tagController.InitTagController(service)

	router.Get("/tag", controller.HandleFindAll)
	router.Get("/tag/:id", controller.HandleFindByID)
	router.Post("/tag", controller.HandleCreateTag)
	router.Patch("/tag/:id", controller.HandleUpdateTag)
	router.Delete("/tag/:id", controller.HandleDeleteTag)
}
