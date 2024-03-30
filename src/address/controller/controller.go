package addressController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"

	addressService "github.com/chronicler-org/core/src/address"
	addressDTO "github.com/chronicler-org/core/src/address/dto/service"
)

type AddressController struct {
	service *addressService.AddressService
}

func InitAddressController(s *addressService.AddressService) *AddressController {
	return &AddressController{
		service: s,
	}
}

func (controller *AddressController) HandleFindByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	address, err := controller.service.FindByID(id)
	return appUtil.PaginateSingle(address), err
}

func (controller *AddressController) HandleCreateAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createAddressDTO addressDTO.CreateAddressDTO

	c.BodyParser(&createAddressDTO)

	addressCreated, err := controller.service.Create(createAddressDTO)

	return appUtil.PaginateSingle(addressCreated), err
}

func (controller *AddressController) HandleUpdateAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateAddressDTO addressDTO.UpdateAddressDTO
	c.BodyParser(&updateAddressDTO)

	id := c.Params("id")

	addressUpdated, err := controller.service.Update(id, updateAddressDTO)

	return appUtil.PaginateSingle(addressUpdated), err
}

func (controller *AddressController) HandleDeleteAddress(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	addressDeleted, err := controller.service.Delete(id)
	return appUtil.PaginateSingle(addressDeleted), err
}
