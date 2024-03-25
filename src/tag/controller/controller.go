package tagController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

type TagController struct {
	service *tagService.TagService
}

func InitTagController(s *tagService.TagService) *TagController {
	return &TagController{
		service: s,
	}
}

func (controller *TagController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDTO appDto.PaginationDTO
	c.QueryParser(&paginationDTO)

	totalCount, tags, err := controller.service.FindAll(paginationDTO)

	return appUtil.Paginate(tags, totalCount, paginationDTO.GetPage(), paginationDTO.GetLimit()), err
}

func (controller *TagController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	tag, err := controller.service.FindByID(id)
	return appUtil.PaginateSingle(tag), err
}

func (controller *TagController) HandleCreateTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createTagDTO tagDTO.CreateTagDTO

	c.BodyParser(&createTagDTO)

	tagCreated, err := controller.service.Create(createTagDTO)

	return appUtil.PaginateSingle(tagCreated), err
}

func (controller *TagController) HandleUpdateTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateTagDTO tagDTO.UpdateTagDTO
	c.BodyParser(&updateTagDTO)

	id := c.Params("id")

	tagUpdated, err := controller.service.Update(id, updateTagDTO)

	return appUtil.PaginateSingle(tagUpdated), err
}

func (controller *TagController) HandleDeleteTag(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	tagDeleted, err := controller.service.Delete(id)
	return appUtil.PaginateSingle(tagDeleted), err
}
