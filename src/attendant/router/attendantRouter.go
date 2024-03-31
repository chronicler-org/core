package attendantRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	attendantController "github.com/chronicler-org/core/src/attendant/controller"
	attendantDTO "github.com/chronicler-org/core/src/attendant/dto"
	attendantRepository "github.com/chronicler-org/core/src/attendant/repository"
	attendantService "github.com/chronicler-org/core/src/attendant/service"
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
	router.Group("/attendant")

	router.Get("/evaluation",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(attendantController.HandleFindAllAttendantEvaluations),
	)
	router.Get("/evaluation/:id",
		appUtil.Controller(attendantController.HandleFindAttendantEvaluationByID),
	)
	router.Post("/evaluation",
		validatorMiddleware(&attendantDTO.CreateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleCreateAttendantEvaluation),
	)
	router.Patch("/evaluation/:id",
		validatorMiddleware(&attendantDTO.UpdateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleUpdateAttendantEvaluation),
	)
	router.Delete("/evaluation/:id",
		appUtil.Controller(attendantController.HandleDeleteAttendantEvaluation),
	)

	router.Get("",
		validatorMiddleware(nil, &appDto.PaginationDTO{}),
		appUtil.Controller(attendantController.HandleFindAllAttendants),
	)
	router.Get("/:id",
		appUtil.Controller(attendantController.HandleFindAttendantByID),
	)
	router.Post("",
		validatorMiddleware(&attendantDTO.CreateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleCreateAttendant),
	)
	router.Patch("/:id",
		validatorMiddleware(&attendantDTO.UpdateAttendantDTO{}, nil),
		appUtil.Controller(attendantController.HandleUpdateAttendant),
	)
	router.Delete("/:id",
		appUtil.Controller(attendantController.HandleDeleteAttendant),
	)
}
