package customerController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	
	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerService "github.com/chronicler-org/core/src/customer/service"
)

type CustomerController struct {
	customerService *customerService.CustomerService
}

func InitCustomerController(s *customerService.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: s,
	}
}

func (controller *CustomerController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, customers, err := controller.customerService.FindAll(paginationDto)

	return appUtil.Paginate(customers, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *CustomerController) HandleFindByCPF(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	cpf := c.Params("cpf")

	customer, err := controller.customerService.FindByCPF(cpf)
	return appUtil.PaginateSingle(customer), err
}

func (controller *CustomerController) HandleCreateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createCustomerDTO customerDTO.CreateCustomerDTO

	c.BodyParser(&createCustomerDTO)

	customerCreated, err := controller.customerService.Create(createCustomerDTO)

	return appUtil.PaginateSingle(customerCreated), err
}

func (controller *CustomerController) HandleUpdateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateCustomerDTO customerDTO.UpdateCustomerDTO
	c.BodyParser(&updateCustomerDTO)

	cpf := c.Params("cpf")

	customerUpdated, err := controller.customerService.Update(cpf, updateCustomerDTO)

	return appUtil.PaginateSingle(customerUpdated), err
}

func (controller *CustomerController) HandleDeleteCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	cpf := c.Params("cpf")

	customerDeleted, err := controller.customerService.Delete(cpf)
	return appUtil.PaginateSingle(customerDeleted), err
}
