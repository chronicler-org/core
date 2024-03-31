package tagRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	tagController "github.com/chronicler-org/core/src/tag/controller"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagRepository "github.com/chronicler-org/core/src/tag/repository"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

func InitTagModule(
	db *gorm.DB,
) (*tagController.TagController, *tagService.TagService) {
	tagRepo := tagRepository.InitTagRepository(db)
	tagServ := tagService.InitTagService(tagRepo)
	tagCtrl := tagController.InitTagController(tagServ)

	return tagCtrl, tagServ
}

func InitTagRouter(router *fiber.App, tagController *tagController.TagController) {
	router.Get("/tag", appMiddleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(tagController.HandleFindAll))
	router.Get("/tag/:id", appUtil.Controller(tagController.HandleFindByID))
	router.Post("/tag", appMiddleware.Validate(&tagDTO.CreateTagDTO{}, nil), appUtil.Controller(tagController.HandleCreateTag))
	router.Patch("/tag/:id", appMiddleware.Validate(&tagDTO.UpdateTagDTO{}, nil), appUtil.Controller(tagController.HandleUpdateTag))
	router.Delete("/tag/:id", appUtil.Controller(tagController.HandleDeleteTag))
}
