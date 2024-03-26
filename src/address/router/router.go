package addressRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appDto "github.com/chronicler-org/core/src/app/dto"
	"github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"

	addressController "github.com/chronicler-org/core/src/address/controller"
	addressDTO "github.com/chronicler-org/core/src/address/dto"
	addressRepository "github.com/chronicler-org/core/src/address/repository"
	addressService "github.com/chronicler-org/core/src/address/service"
)

func InitAddressRouter(router *fiber.App, db *gorm.DB) {

	repository := addressRepository.InitAddressRepository(db)
	service := addressService.InitAddressService(repository)
	controller := addressController.InitAddressController(service)

	router.Get("/address/:id", middleware.Validate(nil, &appDto.PaginationDTO{}), appUtil.Controller(controller.HandleFindByID))
	router.Get("/address/:id", appUtil.Controller(controller.HandleFindByID))
	router.Post("/address", middleware.Validate(&addressDTO.CreateAddressDTO{}, nil), appUtil.Controller(controller.HandleCreateAddress))
	router.Patch("/address/:id", middleware.Validate(&addressDTO.UpdateAddressDTO{}, nil), appUtil.Controller(controller.HandleUpdateAddress))
	router.Delete("/address/:id", appUtil.Controller(controller.HandleDeleteAddress))
}
