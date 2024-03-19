package tagController

import (
	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagService "github.com/chronicler-org/core/src/tag/service"
	"github.com/gofiber/fiber/v2"
)

type TagController struct {
	service *tagService.TagService
}

func InitTagController(s *tagService.TagService) *TagController {
	return &TagController{
		service: s,
	}
}

func (controller *TagController) HandleFindAll(c *fiber.Ctx) error {
	tags, err := controller.service.FindAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(tags)
}

func (controller *TagController) HandleFindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	tag, err := controller.service.FindByID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(tag)
}

func (controller *TagController) HandleCreateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.CreateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newTagID, err := controller.service.Create(tagDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"location": newTagID,
	})
}

func (controller *TagController) HandleUpdateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.UpdateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := c.Params("id")

	tagUpdated, err := controller.service.Update(id, tagDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(tagUpdated)
}

func (controller *TagController) HandleDeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.Delete(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
