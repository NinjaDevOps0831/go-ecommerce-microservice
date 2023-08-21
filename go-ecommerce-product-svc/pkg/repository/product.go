package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/common"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/model/response"
	interfaces "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type productDatabase struct {
	DB *gorm.DB
}

func NewProductRepository(DB *gorm.DB) interfaces.ProductRepository {
	return &productDatabase{DB}
}

func (c *productDatabase) CreateCategory(ctx context.Context, newCategory string) (domain.ProductCategory, error) {
	var createdCategory domain.ProductCategory
	createCategoryQuery := `INSERT INTO product_categories(category_name)
							VALUES($1)
							RETURNING id, category_name` //By including the RETURNING clause, the INSERT statement will not only insert the new row into the table but also return the specified columns as a result. This can be useful when you need to retrieve the generated values or verify the inserted data.
	err := c.DB.Raw(createCategoryQuery, newCategory).Scan(&createdCategory).Error
	if err != nil {
		return domain.ProductCategory{}, err
	}
	return createdCategory, nil
}

//---------brand management------------------

func (c *productDatabase) CreateBrand(ctx context.Context, newBrandDetails domain.ProductBrand) (domain.ProductBrand, error) {
	var createdBrand domain.ProductBrand
	createBrandQuery := `INSERT INTO product_brands (brand_name)
						VALUES($1)
						RETURNING *;`

	err := c.DB.Raw(createBrandQuery, newBrandDetails.BrandName).Scan(&createdBrand).Error
	if err != nil {
		return domain.ProductBrand{}, err
	}
	return createdBrand, nil

}

//---------product management----------------------

func (c *productDatabase) CreateProduct(ctx context.Context, newProduct domain.Product) (domain.Product, error) {
	var createdProduct domain.Product
	productCreateQuery := `INSERT INTO products(product_category_id, name, brand_id, description, product_image)
							VALUES($1,$2,$3,$4,$5)
							RETURNING *`
	err := c.DB.Raw(productCreateQuery, newProduct.ProductCategoryID, newProduct.Name, newProduct.BrandID, newProduct.Description, newProduct.ProductImage).Scan(&createdProduct).Error
	if err != nil {
		return domain.Product{}, err
	}
	return createdProduct, nil
}

func (c *productDatabase) ListAllProducts(ctx context.Context, viewProductsQueryParam common.QueryParams) ([]response.ViewProduct, error) {

	//findQuery := "SELECT * FROM products"
	findQuery := `	SELECT pd.id AS product_details_id, p.name, pb.brand_name, pd.model_no,pd.price, p.description,  p.product_image
					FROM products p
					LEFT JOIN product_details AS pd ON pd.product_id = p.id
					LEFT JOIN product_brands AS pb ON p.brand_id = pb.id`
	params := []interface{}{}

	if viewProductsQueryParam.Query != "" && viewProductsQueryParam.Filter != "" {
		findQuery = fmt.Sprintf("%s WHERE LOWER(%s) LIKE $%d", findQuery, viewProductsQueryParam.Filter, len(params)+1)
		params = append(params, "%"+strings.ToLower(viewProductsQueryParam.Query)+"%")
		fmt.Println("params is ", params)
	}
	if viewProductsQueryParam.SortBy != "" {
		findQuery = fmt.Sprintf("%s ORDER BY %s %s", findQuery, viewProductsQueryParam.SortBy, orderByDirection(viewProductsQueryParam.SortDesc))
	}
	if viewProductsQueryParam.Limit != 0 && viewProductsQueryParam.Page != 0 {
		findQuery = fmt.Sprintf("%s LIMIT $%d OFFSET $%d", findQuery, len(params)+1, len(params)+2)
		params = append(params, viewProductsQueryParam.Limit, (viewProductsQueryParam.Page-1)*viewProductsQueryParam.Limit)
	}

	var allProducts []response.ViewProduct
	err := c.DB.Raw(findQuery, params...).Scan(&allProducts).Error

	if err != nil {
		return nil, err
	}

	return allProducts, err
}

func (c *productDatabase) AddProductDetails(ctx context.Context, NewProductDetails request.NewProductDetails) (domain.ProductDetails, error) {
	var addedProdDetails domain.ProductDetails
	addProdDetailsQuery := `INSERT INTO product_details(product_id,model_no,processor,storage,ram,graphics_card,display_size,color,os,sku,qty_in_stock,price,product_details_image)
							VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
							RETURNING *`

	err := c.DB.Raw(addProdDetailsQuery, NewProductDetails.ProductID, NewProductDetails.ModelNo, NewProductDetails.Processor, NewProductDetails.Storage, NewProductDetails.Ram, NewProductDetails.GraphicsCard, NewProductDetails.DisplaySize, NewProductDetails.Color, NewProductDetails.OS, NewProductDetails.SKU, NewProductDetails.QtyInStock, NewProductDetails.Price, NewProductDetails.ProductDetailsImage).Scan(&addedProdDetails).Error
	if err != nil {
		return domain.ProductDetails{}, err
	}
	return addedProdDetails, nil
}

func orderByDirection(sortDesc bool) string {
	if sortDesc {
		return "DESC"
	}
	return "ASC"
}
