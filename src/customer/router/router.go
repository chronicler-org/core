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
)

func InitCustomerRouter(router *fiber.App, db *gorm.DB) {
	repository := customerRepository.InitCustomerRepository(db)
	service := customerService.InitCustomerService(repository)
	controller := customerController.InitCustomerController(service)

	router.Get("/customer", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(controller.HandleFindAll))
	router.Get("/customer/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/customer", middleware.Validate(&customerDTO.CreateCustomerDTO{}, nil), appUtil.Controller(controller.HandleCreateCustomer))
	router.Patch("/customer/:id", middleware.Validate(&customerDTO.UpdateCustomerDTO{}, nil), appUtil.Controller(controller.HandleUpdateCustomer))
	router.Delete("/customer/:id", appUtil.Controller(controller.HandleDeleteCustomer))
}
