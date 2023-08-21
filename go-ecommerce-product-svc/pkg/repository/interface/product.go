package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/response"
)

type ProductRepository interface {
	CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error)
	//ListAllCategories(ctx context.Context) ([]domain.ProductCategory, error)
	//FindCategoryByID(ctx context.Context, categoryID int) (domain.ProductCategory, error)
	//UpdateCategory(ctx context.Context, updateCatInfo domain.ProductCategory) (domain.ProductCategory, error)
	//DeleteCategory(ctx context.Context, categoryID int) (domain.ProductCategory, error)

	CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error)
	ListAllProducts(ctx context.Context, viewProductsQueryParam common.QueryParams) ([]response.ViewProduct, error)
	//FindProductByID(ctx context.Context, productID int) (domain.Product, error)
	//UpdateProduct(ctx context.Context, updateProductInfo domain.Product) (domain.Product, error)
	//DeleteProduct(ctx context.Context, productID int) error

	CreateBrand(ctx context.Context, newBrandDetails domain.ProductBrand) (domain.ProductBrand, error)

	AddProductDetails(ctx context.Context, NewProductDetails request.NewProductDetails) (domain.ProductDetails, error)
}
