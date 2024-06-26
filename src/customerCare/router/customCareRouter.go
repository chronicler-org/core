package customerCareRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	customerService "github.com/chronicler-org/core/src/customer/service"
	customerCareController "github.com/chronicler-org/core/src/customerCare/controller"
	customerCareDTO "github.com/chronicler-org/core/src/customerCare/dto"
	customerCareRepository "github.com/chronicler-org/core/src/customerCare/repository"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	teamService "github.com/chronicler-org/core/src/team/service"
)

func InitCustomerCareModule(
	db *gorm.DB,
	customerServ *customerService.CustomerService,
	teamServ *teamService.TeamService,
) (*customerCareController.CustomerCareController, *customerCareService.CustomerCareService) {
	customerCareEvaluationRepo := customerCareRepository.InitCustomerCareEvaluationRepository(db)
	customerCareRepo := customerCareRepository.InitCustomerCareRepository(db)
	customerCareServ := customerCareService.InitCustomerCareService(customerCareRepo, customerCareEvaluationRepo, customerServ, teamServ)
	customerCareCtrl := customerCareController.InitCustomerCareController(customerCareServ)

	return customerCareCtrl, customerCareServ
}

func InitCustomerCareRouter(
	router *fiber.App,
	customerCareController *customerCareController.CustomerCareController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	customerCareRouter := router.Group("/customer-care")
	managerAccessMiddleware := appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole})

	customerCareRouter.Get("/evaluation",
		validatorMiddleware(nil, &customerCareDTO.QueryCustomerCareEvaluationDTO{}),
		appUtil.Controller(customerCareController.HandleFindAllCustomerCareEvaluations),
	)
	customerCareRouter.Get("/:id/evaluation",
		appUtil.Controller(customerCareController.HandleFindCustomerCareEvaluationByID),
	)
	customerCareRouter.Post("/:id/evaluation",
		validatorMiddleware(&customerCareDTO.CreateCustomerCareEvaluationDTO{}, nil),
		appUtil.Controller(customerCareController.HandleCreateCustomerCareEvaluation),
	)
	customerCareRouter.Patch("/:id/evaluation",
		managerAccessMiddleware,
		validatorMiddleware(&customerCareDTO.UpdateCustomerCareeEvaluationDTO{}, nil),
		appUtil.Controller(customerCareController.HandleUpdateCustomerCareEvaluation),
	)
	customerCareRouter.Delete("/:id/evaluation",
		managerAccessMiddleware,
		appUtil.Controller(customerCareController.HandleDeleteCustomerCareEvaluation),
	)

	customerCareRouter.Get("/",
		validatorMiddleware(nil, &customerCareDTO.QueryCustomerCareDTO{}),
		appUtil.Controller(customerCareController.HandleFindAllCustomerCares),
	)
	customerCareRouter.Get("/:id",
		appUtil.Controller(customerCareController.HandleFindCustomerCareByID),
	)
	customerCareRouter.Post("/",
		validatorMiddleware(&customerCareDTO.CreateCustomerCareDTO{}, nil),
		appUtil.Controller(customerCareController.HandleCreateCustomerCare),
	)
	customerCareRouter.Delete("/:id",
		managerAccessMiddleware,
		appUtil.Controller(customerCareController.HandleDeleteCustomerCare),
	)
}
