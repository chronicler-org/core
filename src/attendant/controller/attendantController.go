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
func (controller *AttendantController) HandleFindAllAttendants(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDto appDto.PaginationDTO
	c.QueryParser(&paginationDto)

	totalCount, attendants, err := controller.attendantService.FindAllAttendants(paginationDto)

	return appUtil.Paginate(attendants, totalCount, paginationDto.GetPage(), paginationDto.GetLimit()), err
}

func (controller *AttendantController) HandleFindAttendantByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	attendant, err := controller.attendantService.FindAttendantByID(id)
	return appUtil.PaginateSingle(attendant), err
}

func (controller *AttendantController) HandleCreateAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createAttendantDTO attendantDTO.CreateAttendantDTO

	c.BodyParser(&createAttendantDTO)

	attendantCreated, err := controller.attendantService.CreateAttendant(createAttendantDTO)

	return appUtil.PaginateSingle(attendantCreated), err
}

func (controller *AttendantController) HandleUpdateAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updatedAttendantDTO attendantDTO.UpdateAttendantDTO
	c.BodyParser(&updatedAttendantDTO)

	id := c.Params("id")

	attendantUpdated, err := controller.attendantService.UpdateAttendant(id, updatedAttendantDTO)

	return appUtil.PaginateSingle(attendantUpdated), err
}

func (controller *AttendantController) HandleDeleteAttendant(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	attendantDeleted, err := controller.attendantService.DeleteAttendant(id)
	return appUtil.PaginateSingle(attendantDeleted), err
}
