package managerController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerModel "github.com/chronicler-org/core/src/manager/model"
	managerService "github.com/chronicler-org/core/src/manager/service"
)

type ManagerController struct {
	managerService *managerService.ManagerService
}

func InitManagerController(s *managerService.ManagerService) *ManagerController {
	return &ManagerController{
		managerService: s,
	}
}
func (controller *ManagerController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryManagerDTO managerDTO.QueryManagerDTO
	c.QueryParser(&queryManagerDTO)

	totalCount, managers, err := controller.managerService.FindAll(queryManagerDTO)

	return appUtil.Paginate(managers, totalCount, queryManagerDTO.GetPage(), queryManagerDTO.GetLimit()), err
}

func (controller *ManagerController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	manager, err := controller.managerService.FindByID(id)
	return appUtil.PaginateSingle(manager), err
}

func (controller *ManagerController) HandleGetLoggedManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	loggedManager := c.Locals(authEnum.ManagerRole).(managerModel.Manager)

	return appUtil.PaginateSingle(loggedManager), nil
}

func (controller *ManagerController) HandleCreateManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createManagerDTO managerDTO.CreateManagerDTO

	c.BodyParser(&createManagerDTO)

	managerCreated, err := controller.managerService.Create(createManagerDTO)

	return appUtil.PaginateSingle(managerCreated), err
}

func (controller *ManagerController) HandleUpdateManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateManagerDTO managerDTO.UpdateManagerDTO
	c.BodyParser(&updateManagerDTO)

	id := c.Params("id")

	managerUpdated, err := controller.managerService.Update(id, updateManagerDTO)

	return appUtil.PaginateSingle(managerUpdated), err
}

func (controller *ManagerController) HandleDeleteManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	managerDeleted, err := controller.managerService.Delete(id)
	return appUtil.PaginateSingle(managerDeleted), err
}
