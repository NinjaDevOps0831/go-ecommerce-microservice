package request

//model structs input

type NewCategory struct {
	CategoryName string `json:"category_name"`
}

type NewProductDetails struct {
	ProductID           uint    `gorm:"not null" json:"product_id" validate:"required"`
	ModelNo             string  `gorm:"not null" json:"model_no" validate:"required"`
	Processor           string  `gorm:"not null" json:"processor" validate:"required"`
	Storage             string  `gorm:"not null" json:"storage" validate:"required"`
	Ram                 string  ` json:"ram" `
	GraphicsCard        string  `json:"graphics_card" `
	DisplaySize         string  `gorm:"not null" json:"display_size" validate:"required"`
	Color               string  `gorm:"not null" json:"color" validate:"required"`
	OS                  string  `json:"os" validate:"required"`
	SKU                 string  `gorm:"not null" json:"sku" validate:"required"`
	QtyInStock          int     `gorm:"not null" json:"qty_in_stock" validate:"required"`
	Price               float64 `gorm:"not null" json:"price" validate:"required"`
	ProductDetailsImage string  `json:"product_item_image"`
}
