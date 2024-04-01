package tagRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
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

func InitTagRouter(
	router *fiber.App,
	tagController *tagController.TagController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	tagRouter := router.Group("/tag")

	tagRouter.Get("/",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(tagController.HandleFindAll),
	)
	tagRouter.Get("/:id",
		appUtil.Controller(tagController.HandleFindByID),
	)

	tagRouter.Use(appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole}))

	tagRouter.Post("/",
		validatorMiddleware(&tagDTO.CreateTagDTO{}, nil),
		appUtil.Controller(tagController.HandleCreateTag),
	)
	tagRouter.Patch("/:id",
		validatorMiddleware(&tagDTO.UpdateTagDTO{}, nil),
		appUtil.Controller(tagController.HandleUpdateTag),
	)
	tagRouter.Delete("/:id",
		appUtil.Controller(tagController.HandleDeleteTag),
	)
}
