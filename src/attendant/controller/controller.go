package attendantController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantDTO "github.com/chronicler-org/core/src/attendant/dto"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
)

type AttendantController struct {
	attendantService *attendantService.AttendantService
}

func InitAttendantController(s *attendantService.AttendantService) *AttendantController {
	return &AttendantController{
		attendantService: s,
	}
}
func (controller *AttendantController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, attendants, err := controller.attendantService.FindAll(paginationDto)

	return appUtil.Paginate(attendants, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *AttendantController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	attendant, err := controller.attendantService.FindByID(id)
	return appUtil.PaginateSingle(attendant), err
}

func (controller *AttendantController) HandleCreateAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createAttendantDTO attendantDTO.CreateAttendantDTO

	c.BodyParser(&createAttendantDTO)

	attendantCreated, err := controller.attendantService.Create(createAttendantDTO)

	return appUtil.PaginateSingle(attendantCreated), err
}

func (controller *AttendantController) HandleUpdateAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updatedAttendantDTO attendantDTO.UpdateAttendantDTO
	c.BodyParser(&updatedAttendantDTO)

	id := c.Params("id")

	attendantUpdated, err := controller.attendantService.Update(id, updatedAttendantDTO)

	return appUtil.PaginateSingle(attendantUpdated), err
}

func (controller *AttendantController) HandleDeleteAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	attendantDeleted, err := controller.attendantService.Delete(id)
	return appUtil.PaginateSingle(attendantDeleted), err
}
