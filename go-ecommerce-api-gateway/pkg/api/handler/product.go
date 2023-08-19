package handler

import (
	"net/http"
	"strconv"

	client "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/client/interfaces"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/response"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	client client.ProductClient
}

func NewProductHandler(client client.ProductClient) ProductHandler {
	return ProductHandler{
		client: client,
	}
}

// ----------Category Management

// CreateCategory
// @Summary Create new product category
// @ID create-category
// @Description Admins can create new categories from the admin panel
// @Tags Product Category
// @Accept json
// @Produce json
// @Param category_name body request.NewCategory true "New category name"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/categories [post]
func (cr *ProductHandler) CreateCategory(c *gin.Context) {
	var category request.NewCategory
	if err := c.Bind(&category); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to process the request", err.Error(), nil))
		return
	}
	//  call the createcategory usecase to create a new category
	createdCategory, err := cr.client.CreateCategory(c.Request.Context(), category)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(400, "failed to create new category", err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(201, "Category Created Succesfully", createdCategory))

}

// ----------Product Brand Management

// CreateBrand
// @Summary Admin can create new product brand
// @ID create-brand
// @Description Admins can create new brands from the admin panel
// @Tags Product Brand
// @Accept json
// @Produce json
// @Param brand_name body domain.ProductBrand true "New brand name"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/brands [post]
func (cr *ProductHandler) CreateBrand(c *gin.Context) {
	var newBrandDetails domain.ProductBrand
	if err := c.Bind(&newBrandDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to process the request", err.Error(), nil))
		return
	}
	//  call the createbrand usecase to create a new category
	createdBrand, err := cr.client.CreateBrand(c.Request.Context(), newBrandDetails)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(400, "failed to create new brand", err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(201, "Brand Created Succesfully", createdBrand))

}

//------Product Management -----------

// product management
// CreateProduct
// @Summary Admin can create new product listings
// @ID create-product
// @Description Admins can create new product listings
// @Tags Products
// @Accept json
// @Produce json
// @Param new_product_details body domain.Product true "new product details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/products/ [post]
func (cr *ProductHandler) CreateProduct(c *gin.Context) {
	var newProduct domain.Product
	if err := c.Bind(&newProduct); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to process the request", err.Error(), nil))
		return
	}
	//  call the createcategory usecase to create a new category
	createdProduct, err := cr.client.CreateProduct(c.Request.Context(), newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to add new product", err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(201, "New product added succesfully", createdProduct))

}

// List All Products
// @Summary List All products
// @Description Admins and users can list all products
// @Tags Products
// @Accept json
// @Produce json
// @Param limit query int false "Number of items to retrieve per page"
// @Param page query int false "Enter the page no to display"
// cpmmenting - query query string false "Search query string"
// commenting - filter query string false "filter criteria for showing the products"
// @Param sort_by query string false "sorting criteria for showing the products"
// @Param sort_desc query bool false "sorting in descending order"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/products [get]
// @Router /admin/products [get]
func (cr *ProductHandler) ListAllProducts(c *gin.Context) {
	var viewProductsQueryParam common.QueryParams

	viewProductsQueryParam.Page, _ = strconv.Atoi(c.Query("page"))
	viewProductsQueryParam.Limit, _ = strconv.Atoi(c.Query("limit"))
	viewProductsQueryParam.Query = c.Query("query")
	viewProductsQueryParam.Filter = c.Query("filter")
	viewProductsQueryParam.SortBy = c.Query("sort_by")
	viewProductsQueryParam.SortDesc, _ = strconv.ParseBool(c.Query("sort_desc"))

	allProducts, err := cr.client.ListAllProducts(c.Request.Context(), viewProductsQueryParam)

	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(500, "failed to fetch the products", err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, response.SuccessResponse(200, "Succesfully fetched all products", allProducts))

}

//--------------PRODUCT DETAILS---------

// AddProductDetails
// @Summary Add a product details
// @ID add-product-details
// @Description This endpoint allows an admin user to add the product details.
// @Tags Product Details
// @Accept json
// @Produce json
// @Param product_details body request.NewProductDetails true "Product details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/product-details/ [post]
func (cr *ProductHandler) AddProductDetails(c *gin.Context) {
	var NewProductDetails request.NewProductDetails
	if err := c.Bind(&NewProductDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to process the request", err.Error(), nil))
		return
	}

	addedProdDetails, err := cr.client.AddProductDetails(c.Request.Context(), NewProductDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to add the product details", err.Error(), nil))
		return
	}
	c.JSON(http.StatusCreated, response.SuccessResponse(201, "Succesfully added the product details", addedProdDetails))

}
