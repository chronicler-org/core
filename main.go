package main

import (
	"log"
	"os"

	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	attendantRouter "github.com/chronicler-org/core/src/attendant/router"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerRouter "github.com/chronicler-org/core/src/customer/router"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerRouter "github.com/chronicler-org/core/src/manager/router"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagRouter "github.com/chronicler-org/core/src/tag/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// inicializa o banco de dados
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	// realiza migration das entidades no banco de dados
	db.AutoMigrate(&managerModel.Manager{}, &customerModel.Customer{}, &tagModel.Tag{}, &attendantModel.Attendant{})

	// inicializa app principal
	app := fiber.New()

	// instancia logger que permite visualizacao das rotas acessadas e status codes retornados
	app.Use(logger.New())

	// rota raiz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 2!")
	})

	// instancia as rotas para cada entidade
	tagService := tagRouter.InitTagRouter(app, db)

	managerRouter.InitManagerRouter(app, db)
	customerRouter.InitCustomerRouter(app, db, tagService)
	attendantRouter.InitAttendantRouter(app, db)

	app.Listen(":8080")
}
