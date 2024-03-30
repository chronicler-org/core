package customerCareRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerService "github.com/chronicler-org/core/src/customer/service"
	customerCareController "github.com/chronicler-org/core/src/customerCare/controller"
	customerCareDTO "github.com/chronicler-org/core/src/customerCare/dto"
	customerCareRepository "github.com/chronicler-org/core/src/customerCare/repository"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	teamService "github.com/chronicler-org/core/src/team/service"
)

func InitCustomerCareRouter(
	router *fiber.App, db *gorm.DB,
	customerServ *customerService.CustomerService,
	teamServ *teamService.TeamService,
) {
	customerCareEvaluationRepository := customerCareRepository.InitCustomerCareEvaluationRepository(db)
	customerCareRepository := customerCareRepository.InitCustomerCareRepository(db)
	customerCareService := customerCareService.InitCustomerCareService(customerCareRepository, customerCareEvaluationRepository, customerServ, teamServ)
	customerCareController := customerCareController.InitCustomerCareController(customerCareService)

	router.Group("/customer-care")
	router.Get("/", middleware.Validate(nil, &customerCareDTO.QueryCustomerCareDTO{}), appUtil.Controller(customerCareController.HandleFindAllCustomerCares))
	router.Get("/:id", appUtil.Controller(customerCareController.HandleFindCustomerCareByID))
	router.Post("/", middleware.Validate(&customerCareDTO.CreateCustomerCareDTO{}, nil), appUtil.Controller(customerCareController.HandleCreateCustomerCare))
	router.Delete("/:id", appUtil.Controller(customerCareController.HandleDeleteCustomerCare))

	router.Get("/evaluation", middleware.Validate(nil, &customerCareDTO.QueryCustomerCareEvaluationDTO{}), appUtil.Controller(customerCareController.HandleFindAllCustomerCareEvaluations))

	router.Get("/:id/evaluation", appUtil.Controller(customerCareController.HandleFindCustomerCareEvaluationByID))
	router.Post("/:id/evaluation", middleware.Validate(&customerCareDTO.CreateCustomerCareEvaluationDTO{}, nil), appUtil.Controller(customerCareController.HandleCreateCustomerCareEvaluation))
	router.Patch("/:id/evaluation", middleware.Validate(&customerCareDTO.UpdateCustomerCareeEvaluationDTO{}, nil), appUtil.Controller(customerCareController.HandleUpdateCustomerCareEvaluation))
	router.Delete("/:id/evaluation", appUtil.Controller(customerCareController.HandleDeleteCustomerCareEvaluation))
}
