package customerController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerService "github.com/chronicler-org/core/src/customer/service"
)

type CustomerController struct {
	service *customerService.CustomerService
}

func InitCustomerController(s *customerService.CustomerService) *CustomerController {
	return &CustomerController{
		service: s,
	}
}

func (controller *CustomerController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, customers, err := controller.service.FindAll(paginationDto)

	return appUtil.Paginate(customers, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *CustomerController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customer, err := controller.service.FindByID(id)
	return appUtil.PaginateSingle(customer), err
}

func (controller *CustomerController) HandleCreateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createCustomerDTO customerDTO.CreateCustomerDTO

	c.BodyParser(&createCustomerDTO)

	customerCreated, err := controller.service.Create(createCustomerDTO)

	return appUtil.PaginateSingle(customerCreated), err
}

func (controller *CustomerController) HandleUpdateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateCustomerDTO customerDTO.UpdateCustomerDTO
	c.BodyParser(&updateCustomerDTO)

	id := c.Params("id")

	customerUpdated, err := controller.service.Update(id, updateCustomerDTO)

	return appUtil.PaginateSingle(customerUpdated), err
}

func (controller *CustomerController) HandleDeleteCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customerDeleted, err := controller.service.Delete(id)
	return appUtil.PaginateSingle(customerDeleted), err
}
