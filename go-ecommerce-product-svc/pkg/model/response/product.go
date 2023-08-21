package response

type ViewProduct struct {
	ProductDetailsID uint    `json:"product_details_id"`
	Name             string  `json:"name"`
	BrandName        string  `json:"brand_name"`
	ModelNo          string  `json:"model_no"`
	Price            float64 `json:"price"`
	Description      string  `json:"description"`
	ProductImage     string  `json:"product_image"`
}
