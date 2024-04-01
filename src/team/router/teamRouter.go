package teamRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
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

func InitTeamRouter(
	router *fiber.App,
	teamController *teamController.TeamController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	teamRouter := router.Group("/team")

	teamRouter.Get("/",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(teamController.HandleFindAll),
	)
	teamRouter.Get("/:id",
		appUtil.Controller(teamController.HandleFindByID),
	)

	teamRouter.Use(appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole}))

	teamRouter.Post("/",
		validatorMiddleware(&teamDTO.CreateTeamDTO{}, nil),
		appUtil.Controller(teamController.HandleCreateTeam),
	)
	teamRouter.Patch("/:id",
		validatorMiddleware(&teamDTO.UpdateTeamDTO{}, nil),
		appUtil.Controller(teamController.HandleUpdateTeam),
	)
	teamRouter.Delete("/:id",
		appUtil.Controller(teamController.HandleDeleteTeam),
	)
}
