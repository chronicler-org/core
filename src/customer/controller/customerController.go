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

func (controller *CustomerController) HandleFindAllCustomers(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, customers, err := controller.customerService.FindAllCustomers(paginationDto)

	return appUtil.Paginate(customers, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *CustomerController) HandleFindCustomerByCPF(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	cpf := c.Params("cpf")

	customer, err := controller.customerService.FindCustomerByCPF(cpf)
	return appUtil.PaginateSingle(customer), err
}

func (controller *CustomerController) HandleCreateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createCustomerDTO customerDTO.CreateCustomerDTO

	c.BodyParser(&createCustomerDTO)

	customerCreated, err := controller.customerService.CreateCustomer(createCustomerDTO)

	return appUtil.PaginateSingle(customerCreated), err
}

func (controller *CustomerController) HandleUpdateCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateCustomerDTO customerDTO.UpdateCustomerDTO
	c.BodyParser(&updateCustomerDTO)

	cpf := c.Params("cpf")

	customerUpdated, err := controller.customerService.UpdateCustomer(cpf, updateCustomerDTO)

	return appUtil.PaginateSingle(customerUpdated), err
}

func (controller *CustomerController) HandleDeleteCustomer(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	cpf := c.Params("cpf")

	customerDeleted, err := controller.customerService.DeleteCustomer(cpf)
	return appUtil.PaginateSingle(customerDeleted), err
}

func (controller *CustomerController) HandleFindCustomerAddressByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customerAddress, err := controller.customerService.FindCustomerAddressByID(id)
	return appUtil.PaginateSingle(customerAddress), err
}

func (controller *CustomerController) HandleFindAllCustomerAddresses(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, customerAddresses, err := controller.customerService.FindAllCustomerAddresses(paginationDto)

	return appUtil.Paginate(customerAddresses, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *CustomerController) HandleCreateCustomerAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createCustomerAddressDTO customerDTO.CreateCustomerAddressDTO

	c.BodyParser(&createCustomerAddressDTO)

	customerAddressCreated, err := controller.customerService.CreateCustomerAddress(createCustomerAddressDTO)
	return appUtil.PaginateSingle(customerAddressCreated), err
}

func (controller *CustomerController) HandleUpdateCustomerAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateCustomerAddressDTO customerDTO.UpdateCustomerAddressDTO
	c.BodyParser(&updateCustomerAddressDTO)

	id := c.Params("id")

	customerAddressUpdated, err := controller.customerService.UpdateCustomerAddress(id, updateCustomerAddressDTO)
	return appUtil.PaginateSingle(customerAddressUpdated), err
}

func (controller *CustomerController) HandleDeleteCustomerAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customerAddressDeleted, err := controller.customerService.DeleteCustomerAddress(id)
	return appUtil.PaginateSingle(customerAddressDeleted), err
}