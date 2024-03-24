package managerController

import (
	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	managerDTO "github.com/chronicler-org/core/src/manager/dto"
	managerService "github.com/chronicler-org/core/src/manager/service"
	"github.com/gofiber/fiber/v2"
)

type ManagerController struct {
	service *managerService.ManagerService
}

func InitManagerController(s *managerService.ManagerService) *ManagerController {
	return &ManagerController{
		service: s,
	}
}
func (controller *ManagerController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, managers, err := controller.service.FindAll(paginationDto)

	return appUtil.Paginate(managers, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *ManagerController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	manager, err := controller.service.FindByID(id)
	return appUtil.PaginateSingle(manager), err
}

func (controller *ManagerController) HandleCreateManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var managerDTO managerDTO.CreateManagerDTO

	c.BodyParser(&managerDTO)

	managerCreated, err := controller.service.Create(managerDTO)

	return appUtil.PaginateSingle(managerCreated), err
}

func (controller *ManagerController) HandleUpdateManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var managerDTO managerDTO.UpdateManagerDTO
	c.BodyParser(&managerDTO)

	id := c.Params("id")

	managerUpdated, err := controller.service.Update(id, managerDTO)

	return appUtil.PaginateSingle(managerUpdated), err
}

func (controller *ManagerController) HandleDeleteManager(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	managerDeleted, err := controller.service.Delete(id)
	return appUtil.PaginateSingle(managerDeleted), err
}
