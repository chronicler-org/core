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
)

func InitAttendantRouter(router *fiber.App, db *gorm.DB) {

	repository := attendantRepository.InitAttendantRepository(db)
	service := attendantService.InitAttendantService(repository)
	controller := attendantController.InitAttendantController(service)

	router.Get("/attendant", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(controller.HandleFindAll))
	router.Get("/attendant/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/attendant", middleware.Validate(&attendantDTO.CreateAttendantDTO{}, nil), appUtil.Controller(controller.HandleCreateAttendant))
	router.Patch("/attendant/:id", middleware.Validate(&attendantDTO.UpdateAttendantDTO{}, nil), appUtil.Controller(controller.HandleUpdateAttendant))
	router.Delete("/attendant/:id", appUtil.Controller(controller.HandleDeleteAttendant))
}
