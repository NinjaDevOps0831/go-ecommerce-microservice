package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/pb"
)

type ProductClient interface {
	CreateCategory(ctx context.Context, category request.NewCategory) (*pb.CreateCategoryResponse, error)
	CreateBrand(ctx context.Context, newBrandDetails domain.ProductBrand) (*pb.CreateBrandResponse, error)
	CreateProduct(ctx context.Context, newProduct domain.Product) (*pb.CreateProductResponse, error)
	ListAllProducts(ctx context.Context, viewProductsQueryParam common.QueryParams) (*pb.ListAllProductsResponse, error)
	AddProductDetails(ctx context.Context, NewProductDetails request.NewProductDetails) (*pb.AddProductDetailsResponse, error)
}
