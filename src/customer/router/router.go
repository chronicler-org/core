package customerRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerController "github.com/chronicler-org/core/src/customer/controller"
	customerDTO "github.com/chronicler-org/core/src/customer/dto"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
	customerService "github.com/chronicler-org/core/src/customer/service"
	tagService "github.com/chronicler-org/core/src/tag/service"
)

func InitCustomerModule(
	db *gorm.DB,
	tagServ *tagService.TagService,
) (*customerController.CustomerController, *customerService.CustomerService) {
	customerRepo := customerRepository.InitCustomerRepository(db)
	customerServ := customerService.InitCustomerService(customerRepo, tagServ)
	customerCtrl := customerController.InitCustomerController(customerServ)

	return customerCtrl, customerServ
}

func InitCustomerRouter(
	router *fiber.App,
	customerController *customerController.CustomerController,
) {
	router.Get("/customer", appMiddleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(customerController.HandleFindAll))
	router.Get("/customer/:cpf", appUtil.Controller(customerController.HandleFindByCPF))
	router.Post("/customer", appMiddleware.Validate(&customerDTO.CreateCustomerDTO{}, nil), appUtil.Controller(customerController.HandleCreateCustomer))
	router.Patch("/customer/:cpf", appMiddleware.Validate(&customerDTO.UpdateCustomerDTO{}, nil), appUtil.Controller(customerController.HandleUpdateCustomer))
	router.Delete("/customer/:cpf", appUtil.Controller(customerController.HandleDeleteCustomer))
}
