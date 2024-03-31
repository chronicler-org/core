package attendantRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
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

func InitAttendantRouter(router *fiber.App, attendantController *attendantController.AttendantController) {
	router.Get("/attendant", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(attendantController.HandleFindAllAttendants))
	router.Get("/attendant/:id", appUtil.Controller(attendantController.HandleFindAttendantByID))
	router.Post("/attendant", middleware.Validate(&attendantDTO.CreateAttendantDTO{}, nil), appUtil.Controller(attendantController.HandleCreateAttendant))
	router.Patch("/attendant/:id", middleware.Validate(&attendantDTO.UpdateAttendantDTO{}, nil), appUtil.Controller(attendantController.HandleUpdateAttendant))
	router.Delete("/attendant/:id", appUtil.Controller(attendantController.HandleDeleteAttendant))
}
