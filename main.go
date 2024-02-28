package main

import (
	customerRouter "github.com/chronicler-org/core/src/customer/router"
	managerRouter "github.com/chronicler-org/core/src/manager/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Mount("/manager", managerRouter.NewManagerRouter())
	app.Mount("/customer", customerRouter.NewCustomerRouter())

	app.Listen(":8080")
}
