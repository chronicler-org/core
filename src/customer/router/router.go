package customerRouter

import (
	customerController "github.com/chronicler-org/core/src/customer/controller"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
	customerService "github.com/chronicler-org/core/src/customer/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitCustomerRouter(router *fiber.App, db *gorm.DB) {
	repository := customerRepository.InitCustomerRepository(db)
	service := customerService.InitCustomerService(repository)
	controller := customerController.InitCustomerController(service)

	router.Get("/customer", controller.HandleFindAll)
	router.Get("/customer/:id", controller.HandleFindByID)
	router.Post("/customer", controller.HandleCreateCustomer)
	router.Patch("/customer/:id", controller.HandleUpdateCustomer)
	router.Delete("/customer/:id", controller.HandleDeleteCustomer)
}
