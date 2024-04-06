package customerRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

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
	customerAddressRepo := customerRepository.InitCustomerAddressRepository(db)
	customerServ := customerService.InitCustomerService(customerRepo, customerAddressRepo, tagServ)
	customerCtrl := customerController.InitCustomerController(customerServ)

	return customerCtrl, customerServ
}

func InitCustomerRouter(
	router *fiber.App,
	customerController *customerController.CustomerController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	customerRouter := router.Group("/customer")

	customerRouter.Get("/address",
		validatorMiddleware(nil, &customerDTO.QueryCustomerAddressDTO{}),
		appUtil.Controller(customerController.HandleFindAllCustomerAddresses),
	)
	customerRouter.Get("/address/:id",
		appUtil.Controller(customerController.HandleFindCustomerAddressByID),
	)
	customerRouter.Post("/address",
		validatorMiddleware(&customerDTO.CreateCustomerAddressDTO{}, nil),
		appUtil.Controller(customerController.HandleCreateCustomerAddress),
	)
	customerRouter.Patch("/address/:id",
		validatorMiddleware(&customerDTO.UpdateCustomerAddressDTO{}, nil),
		appUtil.Controller(customerController.HandleUpdateCustomerAddress),
	)
	customerRouter.Delete("/address/:id",
		appUtil.Controller(customerController.HandleDeleteCustomerAddress),
	)

	customerRouter.Get("/",
		validatorMiddleware(nil, &customerDTO.QueryCustomerDTO{}),
		appUtil.Controller(customerController.HandleFindAllCustomers),
	)
	customerRouter.Get("/new-variation-percent",
		appUtil.Controller(customerController.HandleGetNewCustomersVariationPercent),
	)
	customerRouter.Post("/",
		validatorMiddleware(&customerDTO.CreateCustomerDTO{}, nil),
		appUtil.Controller(customerController.HandleCreateCustomer))
	customerRouter.Get("/:cpf",
		appUtil.Controller(customerController.HandleFindCustomerByCPF),
	)
	customerRouter.Patch("/:cpf",
		validatorMiddleware(&customerDTO.UpdateCustomerDTO{}, nil),
		appUtil.Controller(customerController.HandleUpdateCustomer),
	)
	customerRouter.Delete("/:cpf",
		appUtil.Controller(customerController.HandleDeleteCustomer),
	)

}
