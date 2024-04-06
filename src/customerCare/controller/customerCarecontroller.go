package customerCareController

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantModel "github.com/chronicler-org/core/src/attendant/model"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	customerCareDTO "github.com/chronicler-org/core/src/customerCare/dto"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	managerModel "github.com/chronicler-org/core/src/manager/model"
)

type CustomerCareController struct {
	customerCareService *customerCareService.CustomerCareService
}

func InitCustomerCareController(customerCareService *customerCareService.CustomerCareService) *CustomerCareController {
	return &CustomerCareController{
		customerCareService: customerCareService,
	}
}

func (controller *CustomerCareController) HandleFindAllCustomerCares(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryCustomerCareDTO customerCareDTO.QueryCustomerCareDTO
	c.QueryParser(&queryCustomerCareDTO)

	totalCount, customerCares, err := controller.customerCareService.FindAllCustomerCares(queryCustomerCareDTO)
	return appUtil.Paginate(customerCares, totalCount, queryCustomerCareDTO.GetPage(), queryCustomerCareDTO.GetLimit()), err
}

func (controller *CustomerCareController) HandleFindCustomerCareByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customerCare, err := controller.customerCareService.FindCustomerCareByID(id)
	return appUtil.PaginateSingle(customerCare), err
}

func (controller *CustomerCareController) HandleCreateCustomerCare(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var teamId uuid.UUID
	if c.Locals(authEnum.ManagerRole) != nil {
		loggedManager := c.Locals(authEnum.ManagerRole).(managerModel.Manager)
		teamId = loggedManager.Team.ID
	} else {
		loggedAttendant := c.Locals(authEnum.AttendantRole).(attendantModel.Attendant)
		teamId = loggedAttendant.Team.ID
	}

	var createCustomerCareDTO customerCareDTO.CreateCustomerCareDTO
	c.BodyParser(&createCustomerCareDTO)

	customerCareCreated, err := controller.customerCareService.CreateCustomerCare(createCustomerCareDTO, teamId)
	return appUtil.PaginateSingle(customerCareCreated), err
}

func (controller *CustomerCareController) HandleDeleteCustomerCare(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	customerCareDeleted, err := controller.customerCareService.DeleteCustomerCare(id)
	return appUtil.PaginateSingle(customerCareDeleted), err
}

func (controller *CustomerCareController) HandleFindAllCustomerCareEvaluations(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryCustomerCareEvaluationDTO customerCareDTO.QueryCustomerCareEvaluationDTO
	c.QueryParser(&queryCustomerCareEvaluationDTO)

	totalCount, customerCareEvaluations, err := controller.customerCareService.FindAllCustomerCareEvaluations(queryCustomerCareEvaluationDTO)
	return appUtil.Paginate(customerCareEvaluations, totalCount, queryCustomerCareEvaluationDTO.GetPage(), queryCustomerCareEvaluationDTO.GetLimit()), err
}

func (controller *CustomerCareController) HandleFindCustomerCareEvaluationByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	customerCareId := c.Params("id")

	customerCareEvaluation, err := controller.customerCareService.FindCustomerCareEvaluationByID(customerCareId)
	return appUtil.PaginateSingle(customerCareEvaluation), err
}

func (controller *CustomerCareController) HandleCreateCustomerCareEvaluation(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	customerCareId := c.Params("id")
	var createCustomerCareEvaluationDTO customerCareDTO.CreateCustomerCareEvaluationDTO

	c.BodyParser(&createCustomerCareEvaluationDTO)

	customerCareEvaluationCreated, err := controller.customerCareService.CreateCustomerCareEvaluation(customerCareId, createCustomerCareEvaluationDTO)
	return appUtil.PaginateSingle(customerCareEvaluationCreated), err
}

func (controller *CustomerCareController) HandleUpdateCustomerCareEvaluation(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	customerCareId := c.Params("id")
	var updateCustomerCareeEvaluationDTO customerCareDTO.UpdateCustomerCareeEvaluationDTO
	c.BodyParser(&updateCustomerCareeEvaluationDTO)

	customerCareEvaluationUpdated, err := controller.customerCareService.UpdateCustomerCareEvaluation(customerCareId, updateCustomerCareeEvaluationDTO)

	return appUtil.PaginateSingle(customerCareEvaluationUpdated), err
}

func (controller *CustomerCareController) HandleDeleteCustomerCareEvaluation(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	customerCareId := c.Params("id")

	customerCareEvaluationDeleted, err := controller.customerCareService.DeleteCustomerCareEvaluation(customerCareId)
	return appUtil.PaginateSingle(customerCareEvaluationDeleted), err
}
