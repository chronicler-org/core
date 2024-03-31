package productController

import (
	"github.com/gofiber/fiber/v2"

	appUtil "github.com/chronicler-org/core/src/app/utils"
	productDTO "github.com/chronicler-org/core/src/product/dto"
	productService "github.com/chronicler-org/core/src/product/service"
)

type ProductController struct {
	productService *productService.ProductService
}

func InitManagerController(productService *productService.ProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}
func (controller *ProductController) HandleFindAllProducts(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var queryProductDTO productDTO.QueryProductDTO
	c.QueryParser(&queryProductDTO)

	totalCount, products, err := controller.productService.FindAllProducts(queryProductDTO)

	return appUtil.Paginate(products, totalCount, queryProductDTO.GetPage(), queryProductDTO.GetLimit()), err
}

func (controller *ProductController) HandleFindProductByID(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	product, err := controller.productService.FindProductByID(id)
	return appUtil.PaginateSingle(product), err
}

func (controller *ProductController) HandleCreateProduct(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var createProductDTO productDTO.CreateProductDTO

	c.BodyParser(&createProductDTO)

	productCreated, err := controller.productService.CreateProduct(createProductDTO)
	return appUtil.PaginateSingle(productCreated), err
}

func (controller *ProductController) HandleUpdateProduct(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	var updateProductDTO productDTO.UpdateProductDTO
	c.BodyParser(&updateProductDTO)

	id := c.Params("id")

	productUpdated, err := controller.productService.UpdateProduct(id, updateProductDTO)
	return appUtil.PaginateSingle(productUpdated), err
}

func (controller *ProductController) HandleDeleteProduct(c *fiber.Ctx) (appUtil.PaginateResponse, error) {
	id := c.Params("id")

	productDeleted, err := controller.productService.DeleteProduct(id)
	return appUtil.PaginateSingle(productDeleted), err
}
