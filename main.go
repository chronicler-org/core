package main

import (
	"github.com/chronicler-org/core/src/manager/router"
	tagRouter "github.com/chronicler-org/core/src/tag/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Mount("/managers", managerRouter.NewManagerRouter())
	app.Mount("/tag", tagRouter.NewTagRouter())

	app.Listen(":8080")
}
