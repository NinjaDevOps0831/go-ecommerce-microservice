package usecase

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/response"
	interfaces "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/repository/interface"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/usecase/interface"
)

type productUseCase struct {
	productRepo interfaces.ProductRepository
}

func NewProductUseCase(repo interfaces.ProductRepository) services.ProductUseCase {
	return &productUseCase{
		productRepo: repo,
	}

}

func (c *productUseCase) CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error) {
	createdCategory, err := c.productRepo.CreateCategory(ctx, newCategory)
	if err != nil {
		return domain.ProductCategory{}, err
	}
	return createdCategory, nil
}

// -----------Brand Management -----------
func (c *productUseCase) CreateBrand(ctx context.Context, newBrandDetails domain.ProductBrand) (domain.ProductBrand, error) {
	createdBrand, err := c.productRepo.CreateBrand(ctx, newBrandDetails)
	if err != nil {
		return domain.ProductBrand{}, err
	}
	return createdBrand, nil
}

//----------Product Management

func (c *productUseCase) CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error) {
	createdProduct, err := c.productRepo.CreateProduct(ctx, newProduct)
	if err != nil {
		return domain.Product{}, err
	}
	return createdProduct, nil
}

func (c *productUseCase) ListAllProducts(ctx context.Context, viewProductsQueryParam common.QueryParams) ([]response.ViewProduct, error) {
	allProducts, err := c.productRepo.ListAllProducts(ctx, viewProductsQueryParam)
	if err != nil {
		return nil, err
	}
	return allProducts, nil
}

//------ Product Details manage

func (c *productUseCase) AddProductDetails(ctx context.Context, NewProductDetails request.NewProductDetails) (domain.ProductDetails, error) {
	addedProdDetails, err := c.productRepo.AddProductDetails(ctx, NewProductDetails)
	if err != nil {
		return domain.ProductDetails{}, err
	}
	return addedProdDetails, nil

}
