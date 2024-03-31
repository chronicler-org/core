package main

import (
	"log"
	"os"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	appRouter "github.com/chronicler-org/core/src/app/router"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	customerModel "github.com/chronicler-org/core/src/customer/model"
	customerCareModel "github.com/chronicler-org/core/src/customerCare/model"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	productModel "github.com/chronicler-org/core/src/product/model"
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
	db.AutoMigrate(&managerModel.Manager{}, &customerModel.Customer{}, &tagModel.Tag{}, &attendantModel.Attendant{}, &attendantModel.AttendantEvaluation{}, &teamModel.Team{}, &customerCareModel.CustomerCare{}, &customerCareModel.CustomerCareEvaluation{}, &productModel.Product{})

	// inicializa app principal
	app := fiber.New()

	// instancia logger que permite visualizacao das rotas acessadas e status codes retornados
	app.Use(logger.New())

	// rota raiz
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// validator
	Validator := validator.New()

	pt := pt_BR.New()
	uni := ut.New(pt, pt)
	trans, _ := uni.GetTranslator("pt_BR")

	// Registro da validação de CPF
	appUtil.RegisterCPFValidationAndTranslation(Validator, trans)

	appRouter.InitAppRouter(app, db, Validator)

	app.Listen(":8080")
}
