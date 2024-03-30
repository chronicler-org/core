package teamRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	teamController "github.com/chronicler-org/core/src/team/controller"
	teamDTO "github.com/chronicler-org/core/src/team/dto"
	teamRepository "github.com/chronicler-org/core/src/team/repository"
	teamService "github.com/chronicler-org/core/src/team/service"
)

func InitTeamModule(
	db *gorm.DB,
) (*teamController.TeamController, *teamService.TeamService) {
	teamRepo := teamRepository.InitTeamRepository(db)
	teamServ := teamService.InitTeamService(teamRepo)
	teamCtrl := teamController.InitTeamController(teamServ)

	return teamCtrl, teamServ
}

func InitTeamRouter(router *fiber.App,
	teamController *teamController.TeamController,
) {
	router.Get("/team", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(teamController.HandleFindAll))
	router.Get("/team/:id", appUtil.Controller(teamController.HandleFindByID))
	router.Post("/team", middleware.Validate(&teamDTO.CreateTeamDTO{}, nil), appUtil.Controller(teamController.HandleCreateTeam))
	router.Patch("/team/:id", middleware.Validate(&teamDTO.UpdateTeamDTO{}, nil), appUtil.Controller(teamController.HandleUpdateTeam))
	router.Delete("/team/:id", appUtil.Controller(teamController.HandleDeleteTeam))
}
