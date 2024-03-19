package customerRouter

import (
	customerController "github.com/chronicler-org/core/src/customer/controller"
	customerRepository "github.com/chronicler-org/core/src/customer/repository"
	customerService "github.com/chronicler-org/core/src/customer/service"
	"github.com/gofiber/fiber/v2"
)

func NewCustomerRouter() *fiber.App {
	router := fiber.New()

	repository := customerRepository.InitCustomerRepository()
	service := customerService.InitCustomerService(repository)
	controller := customerController.InitCustomerController(service)

	router.Get("/", controller.HandleFindAll)
	router.Get(":id", controller.HandleFindByID)
	router.Post("/", controller.HandleCreateCustomer)
	router.Patch(":id", controller.HandleUpdateCustomer)
	router.Delete(":id", controller.HandleDeleteCustomer)

	return router
}
