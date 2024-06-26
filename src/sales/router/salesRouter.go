package salesRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	productService "github.com/chronicler-org/core/src/product/service"
	salesController "github.com/chronicler-org/core/src/sales/controller"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	salesRepository "github.com/chronicler-org/core/src/sales/repository"
	saleService "github.com/chronicler-org/core/src/sales/service"
)

func InitSalesModule(
	db *gorm.DB,
	customerCareService *customerCareService.CustomerCareService,
	productService *productService.ProductService,
) (*salesController.SalesController, *saleService.SaleService) {
	saleItemRepository := salesRepository.InitSaleItemRepository(db)
	saleRepository := salesRepository.InitSaleRepository(db)
	saleService := saleService.InitSaleService(saleRepository, saleItemRepository, productService, customerCareService)
	salesController := salesController.InitSalesController(saleService)

	return salesController, saleService
}

func InitSalesRouter(
	router *fiber.App,
	salesController *salesController.SalesController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	salesRouter := router.Group("/sale")

	salesRouter.Get(
		"/item",
		validatorMiddleware(nil, &salesDTO.QuerySaleItemDTO{}),
		appUtil.Controller(salesController.HandleFindAllSaleItems),
	)

	salesRouter.Get(
		"/",
		validatorMiddleware(nil, &salesDTO.QuerySalesDTO{}),
		appUtil.Controller(salesController.HandleFindAllSales),
	)
	salesRouter.Get(
		"/products-summary",
		validatorMiddleware(nil, &salesDTO.QuerySalesProductSummaryDTO{}),
		appUtil.Controller(salesController.HandleGetSaleProductsSummary),
	)
	salesRouter.Get(
		"/product-quantity-sold-variation",
		appUtil.Controller(salesController.HandleGetProductQuantitySoldVariation),
	)
	salesRouter.Get(
		"/total-values-sold",
		validatorMiddleware(nil, &salesDTO.QueryTotalSalesSoldDTO{}),
		appUtil.Controller(salesController.HandleGetTotalValuesSold),
	)
	salesRouter.Get(
		"/total-values-sold-variation",
		appUtil.Controller(salesController.HandleGetTotalValueSoldVariation),
	)
	salesRouter.Get(
		"/:id",
		appUtil.Controller(salesController.HandleFindSaleByID),
	)
	salesRouter.Post(
		"/",
		validatorMiddleware(&salesDTO.CreateSaleDTO{}, nil),
		appUtil.Controller(salesController.HandleCreateSale),
	)
	salesRouter.Patch(
		"/:id",
		validatorMiddleware(&salesDTO.UpdateSaleDTO{}, nil),
		appUtil.Controller(salesController.HandleUpdateSale),
	)
	salesRouter.Delete(
		"/:id",
		appUtil.Controller(salesController.HandleDeleteSale),
	)
}
