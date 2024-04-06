package salesController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	saleService "github.com/chronicler-org/core/src/sales/service"
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

func (controller *SalesController) HandleGetTotalValuesSold(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryTotalSalesSoldDTO salesDTO.QueryTotalSalesSoldDTO
	c.QueryParser(&queryTotalSalesSoldDTO)

	produtsSummary, count, err := controller.salesService.GetTotalValuesSold(queryTotalSalesSoldDTO)
	return appUtil.Paginate(produtsSummary, count, queryTotalSalesSoldDTO.GetPage(), queryTotalSalesSoldDTO.GetLimit()), err
}

func (controller *SalesController) HandleGetSaleProductsSummary(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var querySalesProductSummaryDTO salesDTO.QuerySalesProductSummaryDTO
	c.QueryParser(&querySalesProductSummaryDTO)

	produtsSummary, count, err := controller.salesService.GetSaleProductsSummary(querySalesProductSummaryDTO)
	return appUtil.Paginate(produtsSummary, count, querySalesProductSummaryDTO.GetPage(), querySalesProductSummaryDTO.GetLimit()), err
}

func (controller *SalesController) HandleGetProductQuantitySoldVariation(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	productQuantitySoldVariation, err := controller.salesService.GetProductQuantitySoldVariation()
	return appUtil.PaginateSingle(productQuantitySoldVariation), err
}

func (controller *SalesController) HandleFindAllSaleItems(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var querySaleItemDTO salesDTO.QuerySaleItemDTO
	c.QueryParser(&querySaleItemDTO)

	produtsSummary, count, err := controller.salesService.FindAllSaleItems(querySaleItemDTO)
	return appUtil.Paginate(produtsSummary, count, querySaleItemDTO.GetPage(), querySaleItemDTO.GetLimit()), err
}
