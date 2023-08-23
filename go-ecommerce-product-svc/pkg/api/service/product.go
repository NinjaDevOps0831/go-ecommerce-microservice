package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/pb"

	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/usecase/interface"
)

type productServiceServer struct {
	productUseCase services.ProductUseCase
	pb.UnimplementedProductServiceServer
}

func NewProductServiceServer(usecase services.ProductUseCase) pb.ProductServiceServer {
	return &productServiceServer{
		productUseCase: usecase,
	}
}

func (cr *productServiceServer) CreateCategory(ctx context.Context, req *pb.CreateCategoryRequest) (*pb.CreateCategoryResponse, error) {
	category := request.NewCategory{
		CategoryName: req.GetCategoryName(),
	}
	createdCategory, err := cr.productUseCase.CreateCategory(ctx, category.CategoryName)
	if err != nil {
		return &pb.CreateCategoryResponse{Status: http.StatusUnprocessableEntity}, errors.New("failed to create new category")
	}
	data := &pb.ProductCategoryOutput{
		Id:           uint32(createdCategory.ID),
		CategoryName: createdCategory.CategoryName,
	}
	return &pb.CreateCategoryResponse{Status: http.StatusCreated, Response: "Category Created Succesfullyy", ProductCategoryOutput: data}, nil

}

func (cr *productServiceServer) CreateBrand(ctx context.Context, req *pb.CreateBrandRequest) (*pb.CreateBrandResponse, error) {
	newBrandDetails := domain.ProductBrand{
		BrandName: req.GetBrandName(),
	}

	createdBrand, err := cr.productUseCase.CreateBrand(ctx, newBrandDetails)
	if err != nil {
		return &pb.CreateBrandResponse{Status: http.StatusUnprocessableEntity, Response: "failed to create new brand"}, err

	}

	data := &pb.ProductBrandOutput{
		Id:        uint32(createdBrand.ID),
		BrandName: createdBrand.BrandName,
	}
	return &pb.CreateBrandResponse{Status: http.StatusCreated, Response: "Brand Created Succesfully", ProductBrandOutput: data}, nil

}

func (cr *productServiceServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	newProduct := domain.Product{
		ProductCategoryID: uint(req.GetProductCategoryId()),
		Name:              req.GetName(),
		BrandID:           uint(req.GetBrandId()),
		Description:       req.GetDescription(),
		ProductImage:      req.GetProductImage(),
	}

	createdProduct, err := cr.productUseCase.CreateProduct(ctx, newProduct)
	if err != nil {
		return &pb.CreateProductResponse{Status: http.StatusBadRequest, Response: "failed to add new product"}, err

	}

	data := &pb.CreatedProduct{
		Id:                uint32(createdProduct.ID),
		ProductCategoryId: uint32(createdProduct.ProductCategoryID),
		Name:              createdProduct.Name,
		BrandId:           uint32(createdProduct.BrandID),
		Description:       createdProduct.Description,
		ProductImage:      createdProduct.ProductImage,
	}
	return &pb.CreateProductResponse{Status: http.StatusCreated, Response: "New product added succesfully", CreatedProduct: data}, nil

}

func (cr *productServiceServer) ListAllProducts(ctx context.Context, req *pb.ListAllProductsRequest) (*pb.ListAllProductsResponse, error) {
	viewProductsQueryParam := common.QueryParams{
		Page:     int(req.GetPage()),
		Limit:    int(req.GetLimit()),
		SortBy:   req.GetSortBy(),
		SortDesc: req.GetSortByDesc(),
	}

	allProducts, err := cr.productUseCase.ListAllProducts(ctx, viewProductsQueryParam)
	if err != nil {
		return &pb.ListAllProductsResponse{Status: http.StatusInternalServerError, Response: "failed to fetch the products"}, err

	}

	data := make([]*pb.AllProducts, len(allProducts))
	for i, products := range allProducts {
		data[i] = &pb.AllProducts{
			ProductDetailsId: uint32(products.ProductDetailsID),
			Name:             products.Name,
			BrandName:        products.BrandName,
			ModelNo:          products.ModelNo,
			Price:            float32(products.Price),
			Description:      products.Description,
			ProductImage:     products.ProductImage,
		}
	}

	return &pb.ListAllProductsResponse{Status: http.StatusOK, Response: "Succesfully fetched all products", AllProducts: data}, nil

}

func (cr *productServiceServer) AddProductDetails(ctx context.Context, req *pb.AddProductDetailsRequest) (*pb.AddProductDetailsResponse, error) {
	NewProductDetails := request.NewProductDetails{
		ProductID:           uint(req.GetProductId()),
		ModelNo:             req.GetModelNo(),
		Processor:           req.GetProcessor(),
		Storage:             req.GetStorage(),
		Ram:                 req.GetRam(),
		GraphicsCard:        req.GetGraphicsCard(),
		DisplaySize:         req.GetDisplaySize(),
		Color:               req.GetColor(),
		OS:                  req.GetOs(),
		SKU:                 req.GetSku(),
		QtyInStock:          int(req.GetQtyInStock()),
		Price:               req.GetPrice(),
		ProductDetailsImage: req.GetProductItemImage(),
	}

	addedProdDetails, err := cr.productUseCase.AddProductDetails(ctx, NewProductDetails)
	if err != nil {
		return &pb.AddProductDetailsResponse{Status: http.StatusBadRequest, Response: "ffailed to add the product details"}, err

	}

	data := &pb.AddedProdDetails{
		Id:                  uint32(addedProdDetails.ID),
		ProductId:           uint32(addedProdDetails.ProductID),
		ModelNo:             addedProdDetails.ModelNo,
		Processor:           addedProdDetails.Processor,
		Storage:             addedProdDetails.Storage,
		Ram:                 addedProdDetails.Ram,
		GraphicsCard:        addedProdDetails.GraphicsCard,
		DisplaySize:         addedProdDetails.DisplaySize,
		Color:               addedProdDetails.Color,
		Os:                  addedProdDetails.OS,
		Sku:                 addedProdDetails.SKU,
		QtyInStock:          int32(addedProdDetails.QtyInStock),
		Price:               addedProdDetails.Price,
		ProductDetailsImage: addedProdDetails.ProductDetailsImage,
	}
	return &pb.AddProductDetailsResponse{Status: http.StatusCreated, Response: "Succesfully added the product details", AddedProductDetails: data}, nil

}
