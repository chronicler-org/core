package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	appRouter "github.com/chronicler-org/core/src/app/router"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	teamModel "github.com/chronicler-org/core/src/team/model"
)

func main() {
	// inicializa o banco de dados
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	// realiza migration das entidades no banco de dados
	db.AutoMigrate(&managerModel.Manager{}, &customerModel.Customer{}, &tagModel.Tag{}, &attendantModel.Attendant{}, &teamModel.Team{}, &customerCareModel.CustomerCare{}, &customerCareModel.CustomerCareEvaluation{})

	// inicializa app principal
	app := fiber.New()

	// instancia logger que permite visualizacao das rotas acessadas e status codes retornados
	app.Use(logger.New())

	// rota raiz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	appRouter.InitAppRouter(app, db)

	app.Listen(":8080")
}
