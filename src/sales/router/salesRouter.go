package salesRouter

import (
	appUtil "github.com/chronicler-org/core/src/app/utils"
	customerCareService "github.com/chronicler-org/core/src/customerCare/service"
	salesController "github.com/chronicler-org/core/src/sales/controller"
	salesDTO "github.com/chronicler-org/core/src/sales/dto"
	salesRepository "github.com/chronicler-org/core/src/sales/repository"
	saleService "github.com/chronicler-org/core/src/sales/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func InitSalesModule(
	db *gorm.DB,
	customerCareService *customerCareService.CustomerCareService,
) (*salesController.SalesController, *saleService.SaleService) {
	salesRepository := salesRepository.InitSalesRepository(db)
	saleService := saleService.InitSaleService(salesRepository, customerCareService)
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
		"/",
		validatorMiddleware(nil, &salesDTO.QuerySalesDTO{}),
		appUtil.Controller(salesController.HandleFindAllSales),
	)
	salesRouter.Get(
		"/:id",
		appUtil.Controller(salesController.HandleFindSaleByID),
	)
	salesRouter.Post(
		"/:customer_care_id",
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
