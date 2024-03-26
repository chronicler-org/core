package tagController

import (
	"errors"

	tagDTO "github.com/chronicler-org/core/src/tag/dto"
	tagService "github.com/chronicler-org/core/src/tag/service"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(tags)
}

func (controller *TagController) HandleFindByID(c *fiber.Ctx) error {
	id := c.Params("id")
	tag, err := controller.service.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(tag)
}

func (controller *TagController) HandleCreateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.CreateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "erro ao tentar extrair dados do body",
		})
	}

	newTagID, err := controller.service.Create(tagDTO)
	if err != nil {
		target := &serviceErrors.ServiceError{}
		if errors.As(err, &target) {
			c.Status(fiber.StatusBadRequest)
		} else {
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"location": newTagID,
	})
}

func (controller *TagController) HandleUpdateTag(c *fiber.Ctx) error {
	var tagDTO tagDTO.UpdateTagDTO

	err := c.BodyParser(&tagDTO)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "erro ao tentar extrair dados do body",
		})
	}

	id := c.Params("id")

	tagUpdated, err := controller.service.Update(id, tagDTO)
	if err != nil {
		target := &serviceErrors.ServiceError{}
		if errors.As(err, &target) {
			c.Status(fiber.StatusBadRequest)
		} else {
			c.Status(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(tagUpdated)
}

func (controller *TagController) HandleDeleteTag(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusOK)
}
