package customerRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
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
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	customerRouter := router.Group("/customer")

	customerRouter.Get("/",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(customerController.HandleFindAll),
	)
	customerRouter.Get("/:cpf",
		appUtil.Controller(customerController.HandleFindByCPF),
	)
	customerRouter.Post("/",
		validatorMiddleware(&customerDTO.CreateCustomerDTO{}, nil),
		appUtil.Controller(customerController.HandleCreateCustomer))
	customerRouter.Patch("/:cpf",
		validatorMiddleware(&customerDTO.UpdateCustomerDTO{}, nil),
		appUtil.Controller(customerController.HandleUpdateCustomer),
	)
	customerRouter.Delete("/:cpf",
		appUtil.Controller(customerController.HandleDeleteCustomer),
	)
}
