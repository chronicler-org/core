package productRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	authEnum "github.com/chronicler-org/core/src/auth/enum"
	productController "github.com/chronicler-org/core/src/product/controller"
	productDTO "github.com/chronicler-org/core/src/product/dto"
	productRepository "github.com/chronicler-org/core/src/product/repository"
	productService "github.com/chronicler-org/core/src/product/service"
)

func InitProductModule(
	db *gorm.DB,
) (*productController.ProductController, *productService.ProductService) {
	productRepo := productRepository.InitProductRepository(db)
	productServ := productService.InitManagerService(productRepo)
	productCtrl := productController.InitManagerController(productServ)

	return productCtrl, productServ
}

func InitProductRouter(
	router *fiber.App,
	productController *productController.ProductController,
	validatorMiddleware func(interface{}, interface{}) func(*fiber.Ctx) error,
) {
	productRouter := router.Group("/product")

	productRouter.Get("/",
		validatorMiddleware(nil, &productDTO.QueryProductDTO{}),
		appUtil.Controller(productController.HandleFindAllProducts),
	)
	productRouter.Get("/:id",
		appUtil.Controller(productController.HandleFindProductByID),
	)

	productRouter.Use(appMiddleware.RouteAccessMiddleware([]authEnum.Role{authEnum.ManagerRole}))
	
	productRouter.Post("/",
		validatorMiddleware(&productDTO.CreateProductDTO{}, nil),
		appUtil.Controller(productController.HandleCreateProduct),
	)
	productRouter.Patch("/:id",
		validatorMiddleware(&productDTO.UpdateProductDTO{}, nil),
		appUtil.Controller(productController.HandleUpdateProduct),
	)
	productRouter.Delete("/:id",
		appUtil.Controller(productController.HandleDeleteProduct),
	)
}
