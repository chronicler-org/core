package managerController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
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
	var paginationDTO appDto.PaginationDTO
	c.QueryParser(&paginationDTO)

	totalCount, managers, err := controller.managerService.FindAll(paginationDTO)

	return appUtil.Paginate(managers, totalCount, paginationDTO.GetPage(), paginationDTO.GetLimit()), err
}

func (controller *ManagerController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	manager, err := controller.managerService.FindByID(id)
	return appUtil.PaginateSingle(manager), err
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
