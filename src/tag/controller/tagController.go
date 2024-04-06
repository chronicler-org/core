package tagController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

type TagController struct {
	tagService *tagService.TagService
}

func InitTagController(s *tagService.TagService) *TagController {
	return &TagController{
		tagService: s,
	}
}

func (controller *TagController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryTagDTO tagDTO.QueryTagDTO
	c.QueryParser(&queryTagDTO)

	totalCount, tags, err := controller.tagService.FindAll(queryTagDTO)
	return appUtil.Paginate(tags, totalCount, queryTagDTO.GetPage(), queryTagDTO.GetLimit()), err
}

func (controller *TagController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	tag, err := controller.tagService.FindByID(id)
	return appUtil.PaginateSingle(tag), err
}

func (controller *TagController) HandleCreateTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createTagDTO tagDTO.CreateTagDTO

	c.BodyParser(&createTagDTO)

	tagCreated, err := controller.tagService.Create(createTagDTO)

	return appUtil.PaginateSingle(tagCreated), err
}

func (controller *TagController) HandleUpdateTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateTagDTO tagDTO.UpdateTagDTO
	c.BodyParser(&updateTagDTO)

	id := c.Params("id")

	tagUpdated, err := controller.tagService.Update(id, updateTagDTO)

	return appUtil.PaginateSingle(tagUpdated), err
}

func (controller *TagController) HandleDeleteTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	tagDeleted, err := controller.tagService.Delete(id)
	return appUtil.PaginateSingle(tagDeleted), err
}
