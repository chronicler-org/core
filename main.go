package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	attendantRouter "github.com/chronicler-org/core/src/attendant/router"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerRouter "github.com/chronicler-org/core/src/customer/router"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	customerCareRouter "github.com/chronicler-org/core/src/customerCare/router"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerRouter "github.com/chronicler-org/core/src/manager/router"
	tagModel "github.com/chronicler-org/core/src/tag/model"
	tagRouter "github.com/chronicler-org/core/src/tag/router"
	teamModel "github.com/chronicler-org/core/src/team/model"
	teamRouter "github.com/chronicler-org/core/src/team/router"
)

func main() {
	// inicializa o banco de dados
	dbURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	// realiza migration das entidades no banco de dados
	db.AutoMigrate(&managerModel.Manager{}, &customerModel.Customer{}, &tagModel.Tag{}, &attendantModel.Attendant{}, &teamModel.Team{}, &customerCareModel.CustomerCare{})

	// inicializa app principal
	app := fiber.New()

	// instancia logger que permite visualizacao das rotas acessadas e status codes retornados
	app.Use(logger.New())

	// rota raiz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// instancia as rotas para cada entidade
	tagService := tagRouter.InitTagRouter(app, db)
	customerService := customerRouter.InitCustomerRouter(app, db, tagService)
	teamService := teamRouter.InitTeamRouter(app, db)
	managerRouter.InitManagerRouter(app, db, teamService)
	attendantRouter.InitAttendantRouter(app, db, teamService)
	customerCareRouter.InitCustomerCareRouter(app, db, customerService, teamService)

	app.Listen(":8080")
}
