package managerController

import (
	"errors"

	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	"github.com/chronicler-org/core/src/manager/service"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
	"github.com/gofiber/fiber/v2"
)

type ManagerController struct {
	service *managerService.ManagerService
}

func InitManagerController(s *managerService.ManagerService) *ManagerController {
	return &ManagerController{
		service: s,
	}
}
func (controller *ManagerController) HandleFindAll(c *fiber.Ctx) error {
	managers, err := controller.service.FindAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(managers)
}

func (controller *ManagerController) HandleFindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	manager, err := controller.service.FindByID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(manager)
}

func (controller *ManagerController) HandleCreateManager(c *fiber.Ctx) error {
	var managerDTO managerDTO.CreateManagerDTO

	err := c.BodyParser(&managerDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newManagerID, err := controller.service.Create(managerDTO)

	if err != nil {
		target := &serviceErrors.ServiceError{}
		if errors.As(err, &target) {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"location": newManagerID,
	})
}

func (controller *ManagerController) HandleUpdateManager(c *fiber.Ctx) error {
	var managerDTO managerDTO.UpdateManagerDTO

	err := c.BodyParser(&managerDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := c.Params("id")

	managerUpdated, err := controller.service.Update(id, managerDTO)
	if err != nil {
		target := &serviceErrors.ServiceError{}
		if errors.As(err, &target) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(managerUpdated)
}

func (controller *ManagerController) HandleDeleteManager(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.Delete(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
