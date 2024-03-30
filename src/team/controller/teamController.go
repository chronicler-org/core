package teamController

import (
	"github.com/gofiber/fiber/v2"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	teamDTO "github.com/chronicler-org/core/src/team/dto"
	teamService "github.com/chronicler-org/core/src/team/service"
)

type TeamController struct {
	teamService *teamService.TeamService
}

func InitTeamController(s *teamService.TeamService) *TeamController {
	return &TeamController{
		teamService: s,
	}
}

func (controller *TeamController) HandleFindAll(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var paginationDTO appDto.PaginationDTO
	c.QueryParser(&paginationDTO)

	totalCount, teams, err := controller.teamService.FindAll(paginationDTO)

	return appUtil.Paginate(teams, totalCount, paginationDTO.GetPage(), paginationDTO.GetLimit()), err
}

func (controller *TeamController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	team, err := controller.teamService.FindByID(id)
	return appUtil.PaginateSingle(team), err
}

func (controller *TeamController) HandleCreateTeam(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createTeamDTO teamDTO.CreateTeamDTO

	c.BodyParser(&createTeamDTO)

	teamCreated, err := controller.teamService.Create(createTeamDTO)

	return appUtil.PaginateSingle(teamCreated), err
}

func (controller *TeamController) HandleUpdateTeam(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateTeamDTO teamDTO.UpdateTeamDTO
	c.BodyParser(&updateTeamDTO)

	id := c.Params("id")

	teamUpdated, err := controller.teamService.Update(id, updateTeamDTO)

	return appUtil.PaginateSingle(teamUpdated), err
}

func (controller *TeamController) HandleDeleteTeam(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	teamDeleted, err := controller.teamService.Delete(id)
	return appUtil.PaginateSingle(teamDeleted), err
}
