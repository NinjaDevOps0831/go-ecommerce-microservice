package client

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/client/interfaces"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type productClient struct {
	client pb.ProductServiceClient
}

func NewProductClient(cfg *config.Config) (interfaces.ProductClient, error) {

	gcc, err := grpc.Dial(cfg.ProductServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewProductServiceClient(gcc)

	return &productClient{
		client: client,
	}, nil
}

func (cr *productClient) CreateCategory(ctx context.Context, category request.NewCategory) (*pb.CreateCategoryResponse, error) {
	res, err := cr.client.CreateCategory(ctx, &pb.CreateCategoryRequest{
		CategoryName: category.CategoryName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (cr *productClient) CreateBrand(ctx context.Context, newBrandDetails domain.ProductBrand) (*pb.CreateBrandResponse, error) {
	res, err := cr.client.CreateBrand(ctx, &pb.CreateBrandRequest{
		BrandName: newBrandDetails.BrandName,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (cr *productClient) CreateProduct(ctx context.Context, newProduct domain.Product) (*pb.CreateProductResponse, error) {
	res, err := cr.client.CreateProduct(ctx, &pb.CreateProductRequest{
		ProductCategoryId: uint32(newProduct.ProductCategoryID),
		Name:              newProduct.Name,
		BrandId:           uint32(newProduct.BrandID),
		Description:       newProduct.Description,
		ProductImage:      newProduct.ProductImage,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (cr *productClient) ListAllProducts(ctx context.Context, viewProductsQueryParam common.QueryParams) (*pb.ListAllProductsResponse, error) {
	res, err := cr.client.ListAllProducts(ctx, &pb.ListAllProductsRequest{
		Page:       uint32(viewProductsQueryParam.Page),
		Limit:      uint32(viewProductsQueryParam.Limit),
		SortBy:     viewProductsQueryParam.SortBy,
		SortByDesc: viewProductsQueryParam.SortDesc,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}

func (cr *productClient) AddProductDetails(ctx context.Context, NewProductDetails request.NewProductDetails) (*pb.AddProductDetailsResponse, error) {
	res, err := cr.client.AddProductDetails(ctx, &pb.AddProductDetailsRequest{
		ProductId:        uint32(NewProductDetails.ProductID),
		ModelNo:          NewProductDetails.ModelNo,
		Processor:        NewProductDetails.Processor,
		Storage:          NewProductDetails.Storage,
		Ram:              NewProductDetails.Ram,
		GraphicsCard:     NewProductDetails.GraphicsCard,
		DisplaySize:      NewProductDetails.DisplaySize,
		Color:            NewProductDetails.Color,
		Os:               NewProductDetails.OS,
		Sku:              NewProductDetails.SKU,
		QtyInStock:       int32(NewProductDetails.QtyInStock),
		Price:            NewProductDetails.Price,
		ProductItemImage: NewProductDetails.ProductDetailsImage,
	})
	if err != nil {
		return nil, err
	}
	return res, nil

}
