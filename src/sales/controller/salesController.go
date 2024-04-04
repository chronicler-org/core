package salesController

import (
	appUtil "github.com/chronicler-org/core/src/app/utils"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleService "github.com/chronicler-org/core/src/sales/service"
	"github.com/gofiber/fiber/v2"
)

type SalesController struct {
	salesService *saleService.SaleService
}

func InitSalesController(salesService *saleService.SaleService) *SalesController {
	return &SalesController{
		salesService: salesService,
	}
}

func (controller *SalesController) HandleFindAllSales(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var querySalesDTO salesDTO.QuerySalesDTO
	c.QueryParser(&querySalesDTO)

	count, sales, err := controller.salesService.FindAllSales(querySalesDTO)
	return appUtil.Paginate(sales, count, querySalesDTO.GetPage(), querySalesDTO.GetLimit()), err
}

func (controller *SalesController) HandleFindSaleByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")
	sale, err := controller.salesService.FindSaleByID(id)
	return appUtil.PaginateSingle(sale), err
}

func (controller *SalesController) HandleCreateSale(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createSaleDTO salesDTO.CreateSaleDTO
	c.BodyParser(&createSaleDTO)

	createdSale, err := controller.salesService.CreateSale(createSaleDTO)
	return appUtil.PaginateSingle(createdSale), err
}

func (controller *SalesController) HandleUpdateSale(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	customerCareID := c.Params("id")

	var updateSaleDTO salesDTO.UpdateSaleDTO
	c.BodyParser(&updateSaleDTO)

	updateSale, err := controller.salesService.UpdateSale(updateSaleDTO, customerCareID)

	return appUtil.PaginateSingle(updateSale), err
}

func (controller *SalesController) HandleDeleteSale(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	deletedSale, err := controller.salesService.DeleteSale(id)
	return appUtil.PaginateSingle(deletedSale), err
}
