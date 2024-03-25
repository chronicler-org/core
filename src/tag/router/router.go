package tagRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	tagController "github.com/chronicler-org/core/src/tag/controller"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

func InitTagRouter(router *fiber.App, db *gorm.DB) {
	repository := tagRepository.InitTagRepository(db)
	service := tagService.InitTagService(repository)
	controller := tagController.InitTagController(service)

	router.Get("/tag", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(controller.HandleFindAll))
	router.Get("/tag/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/tag", middleware.Validate(&tagDTO.CreateTagDTO{}, nil), appUtil.Controller(controller.HandleCreateTag))
	router.Patch("/tag/:id", middleware.Validate(&tagDTO.UpdateTagDTO{}, nil), appUtil.Controller(controller.HandleUpdateTag))
	router.Delete("/tag/:id", appUtil.Controller(controller.HandleDeleteTag))
}
