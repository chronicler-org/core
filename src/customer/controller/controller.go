package customerController

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerService "github.com/chronicler-org/core/src/customer/service"
	serviceErrors "github.com/chronicler-org/core/src/utils/errors"
)

type CustomerController struct {
	service *customerService.CustomerService
}

func InitCustomerController(s *customerService.CustomerService) *CustomerController {
	return &CustomerController{
		service: s,
	}
}

func (controller *CustomerController) HandleFindAll(c *fiber.Ctx) error {
	customers, err := controller.service.FindAll()
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

func (controller *CustomerController) HandleFindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	customer, err := controller.service.FindByID(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (controller *CustomerController) HandleCreateCustomer(c *fiber.Ctx) error {
	var customerDTO customerDTO.CreateCustomerDTO

	err := c.BodyParser(&customerDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	newCustomerID, err := controller.service.Create(customerDTO)
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
		"location": newCustomerID,
	})
}

func (controller *CustomerController) HandleUpdateCustomer(c *fiber.Ctx) error {
	var customerDTO customerDTO.UpdateCustomerDTO

	err := c.BodyParser(&customerDTO)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	id := c.Params("id")

	customerUpdated, err := controller.service.Update(id, customerDTO)
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

	return c.Status(fiber.StatusOK).JSON(customerUpdated)
}

func (controller *CustomerController) HandleDeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	err := controller.service.Delete(id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.SendStatus(fiber.StatusOK)
}
