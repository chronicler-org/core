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

func InitCustomerRouter(router *fiber.App, db *gorm.DB, tagServ *tagService.TagService) *customerService.CustomerService {
	customerRepository := customerRepository.InitCustomerRepository(db)
	customerService := customerService.InitCustomerService(customerRepository, tagServ)
	customerController := customerController.InitCustomerController(customerService)

	router.Get("/customer", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(customerController.HandleFindAll))
	router.Get("/customer/:cpf", appUtil.Controller(customerController.HandleFindByCPF))
	router.Post("/customer", middleware.Validate(&customerDTO.CreateCustomerDTO{}, nil), appUtil.Controller(customerController.HandleCreateCustomer))
	router.Patch("/customer/:cpf", middleware.Validate(&customerDTO.UpdateCustomerDTO{}, nil), appUtil.Controller(customerController.HandleUpdateCustomer))
	router.Delete("/customer/:cpf", appUtil.Controller(customerController.HandleDeleteCustomer))

	return customerService
}
