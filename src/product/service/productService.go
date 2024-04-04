package productService

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	appException "github.com/chronicler-org/core/src/app/exceptions"
	appUtil "github.com/chronicler-org/core/src/app/utils"
	productDTO "github.com/chronicler-org/core/src/product/dto"
	productEnum "github.com/chronicler-org/core/src/product/enum"
	productExceptionMessage "github.com/chronicler-org/core/src/product/messages"
	productModel "github.com/chronicler-org/core/src/product/model"
	productRepository "github.com/chronicler-org/core/src/product/repository"
)

type ProductService struct {
	productRepository *productRepository.ProductRepository
}

func InitManagerService(productRepository *productRepository.ProductRepository) *ProductService {
	return &ProductService{
		productRepository: productRepository,
	}
}

func (service *ProductService) FindProductByID(id string) (productModel.Product, error) {
	result, err := service.productRepository.FindOneByField("ID", id)
	product, _ := result.(*productModel.Product)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return *product, appException.NotFoundException(productExceptionMessage.PRODUCT_NOT_FOUND)
	}
	return *product, nil
}

func (service *ProductService) CreateProduct(dto productDTO.CreateProductDTO) (productModel.Product, error) {
	model := productModel.Product{
		ID:        uuid.New(),
		Model:     productEnum.ClothingModel(dto.Model),
		Size:      productEnum.Size(dto.Size),
		Value:     dto.Value,
		Fabric:    dto.Fabric,
		Stock:     dto.Stock,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := service.productRepository.Create(model)
	return model, err
}

func (service *ProductService) UpdateProduct(id string, dto productDTO.UpdateProductDTO) (productModel.Product, error) {
	productExists, err := service.FindProductByID(id)
	if err != nil {
		return productModel.Product{}, err
	}
	appUtil.UpdateModelFromDTO(&productExists, dto)

	productExists.UpdatedAt = time.Now()
	err = service.productRepository.Update(productExists)
	return productExists, err
}

func (service *ProductService) FindAllProducts(queryProductDTO productDTO.QueryProductDTO) (int64, []productModel.Product, error) {
	var products []productModel.Product
	totalCount, err := service.productRepository.FindAll(queryProductDTO, &products)
	if err != nil {
		return 0, nil, err
	}
	return totalCount, products, nil
}

func (service *ProductService) DeleteProduct(id string) (productModel.Product, error) {
	productExists, err := service.FindProductByID(id)
	if err != nil {
		return productModel.Product{}, err
	}

	err = service.productRepository.Delete("ID", id)
	if err != nil {
		return productModel.Product{}, err
	}
	return productExists, nil
}

func (service *ProductService) ValidateStock(productID string, requestedQuantity uint32) (productModel.Product, error) {
	product, err := service.FindProductByID(productID)
	if err != nil {
		return productModel.Product{}, err
	}

	if product.Stock < requestedQuantity {
		outOfStock := productExceptionMessage.OUT_OF_STOCK
		outOfStock.Detail = fmt.Sprintf(outOfStock.Detail, product.ID.String(), product.Stock, requestedQuantity)
		return product, appException.ConflictException(outOfStock)
	}

	return product, nil
}

func (service *ProductService) DebitStock(productID string, quantity uint32, tx *gorm.DB) (productModel.Product, error) {
	productHasStockAvailable, err := service.ValidateStock(productID, quantity)
	if err != nil {
		return productModel.Product{}, err
	}

	// Debit the stock
	productHasStockAvailable.Stock -= quantity

	// Update the product with the new stock value
	productHasStockAvailable.UpdatedAt = time.Now()
	err = service.productRepository.UpdateWithTransaction(tx, productHasStockAvailable)
	if err != nil {
		return productModel.Product{}, err
	}

	return productHasStockAvailable, nil
}
