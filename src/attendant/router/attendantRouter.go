package attendantRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantController "github.com/chronicler-org/core/src/attendant/controller"
	attendantDTO "github.com/chronicler-org/core/src/attendant/dto"
	attendantRepository "github.com/chronicler-org/core/src/attendant/repository"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	teamService "github.com/chronicler-org/core/src/team/service"
)

func InitAttendantModule(
	db *gorm.DB,
	teamServ *teamService.TeamService,
) (*attendantController.AttendantController, *attendantService.AttendantService) {
	attendantRepo := attendantRepository.InitAttendantRepository(db)
	attendantEvaluationRepo := attendantRepository.InitAttendantEvaluationRepository(db)
	attendantServ := attendantService.InitAttendantService(attendantRepo, attendantEvaluationRepo, teamServ)
	attendantCtrl := attendantController.InitAttendantController(attendantServ)

	return attendantCtrl, attendantServ
}

func InitAttendantRouter(
	router *fiber.App,
	attendantController *attendantController.AttendantController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	attendantRoute := router.Group("/attendant")
	managerAccessMiddleware := appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole})
	attendantAccessMiddleware := appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.AttendantRole})

	attendantRoute.Get("/evaluation",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(attendantController.HandleFindAllAttendantEvaluations),
	)
	attendantRoute.Get("/evaluation/:id",
		appUtil.Controller(attendantController.HandleFindAttendantEvaluationByID),
	)
	attendantRoute.Post("/evaluation",
		attendantAccessMiddleware,
		validatorMiddleware(&attendantDTO.CreateAttendantEvaluationDTO{}, nil),
		appUtil.Controller(attendantController.HandleCreateAttendantEvaluation),
	)
	attendantRoute.Patch("/evaluation/:id",
		attendantAccessMiddleware,
		validatorMiddleware(&attendantDTO.UpdateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleUpdateAttendantEvaluation),
	)
	attendantRoute.Delete("/evaluation/:id",
		appUtil.Controller(attendantController.HandleDeleteAttendantEvaluation),
	)

	attendantRoute.Get("/",
		validatorMiddleware(nil, &attendantDTO.AttendantQueryDTO{}),
		appUtil.Controller(attendantController.HandleFindAllAttendants),
	)
	attendantRoute.Get("/me",
		attendantAccessMiddleware,
		appUtil.Controller(attendantController.HandleGetLoggedAttendant),
	)
	attendantRoute.Get("/:id",
		appUtil.Controller(attendantController.HandleFindAttendantByID),
	)
	attendantRoute.Post("/",
		managerAccessMiddleware,
		validatorMiddleware(&attendantDTO.CreateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleCreateAttendant),
	)
	attendantRoute.Patch("/:id",
		validatorMiddleware(&attendantDTO.UpdateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleUpdateAttendant),
	)
	attendantRoute.Delete("/:id",
		managerAccessMiddleware,
		appUtil.Controller(attendantController.HandleDeleteAttendant),
	)
}
