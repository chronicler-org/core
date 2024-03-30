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

func InitAttendantRouter(router *fiber.App, db *gorm.DB, teamServ *teamService.TeamService) *attendantService.AttendantService {

	attendantRepository := attendantRepository.InitAttendantRepository(db)
	attendantService := attendantService.InitAttendantService(attendantRepository, teamServ)
	attendantController := attendantController.InitAttendantController(attendantService)

	router.Get("/attendant", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(attendantController.HandleFindAll))
	router.Get("/attendant/:id", appUtil.Controller(attendantController.HandleFindByID))
	router.Post("/attendant", middleware.Validate(&attendantDTO.CreateAttendantDTO{}, nil), appUtil.Controller(attendantController.HandleCreateAttendant))
	router.Patch("/attendant/:id", middleware.Validate(&attendantDTO.UpdateAttendantDTO{}, nil), appUtil.Controller(attendantController.HandleUpdateAttendant))
	router.Delete("/attendant/:id", appUtil.Controller(attendantController.HandleDeleteAttendant))

	return attendantService
}
