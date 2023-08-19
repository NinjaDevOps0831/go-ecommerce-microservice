package domain

type ProductCategory struct {
	ID           uint   `gorm:"primaryKey, uniqueIndex" json:"id"`
	CategoryName string `gorm:"not null, index, unique" json:"category_name"`
}

type Product struct {
	ID                uint            `gorm:"primaryKey,uniqueIndex" json:"id"`
	ProductCategoryID uint            `gorm:"not null" json:"product_category_id" validate:"required"`
	ProductCategory   ProductCategory `gorm:"foreignKey:ProductCategoryID" json:"-"`
	Name              string          `gorm:"not null,uniqueIndex" json:"name" validate:"required"`
	BrandID           uint            `gorm:"not null" json:"brand_id" validate:"required"`
	ProductBrand      ProductBrand    `gorm:"foreignKey:BrandID" json:"-"`
	Description       string          `json:"description"`
	ProductImage      string          `json:"product_image"`
}

type ProductBrand struct {
	ID        uint   `gorm:"primaryKey,uniqueIndex" json:"id"`
	BrandName string `gorm:"not null, index,unique" json:"brand_name"`
}

type ProductDetails struct {
	ID                  uint    `gorm:"primaryKey,uniqueIndex" json:"id"`
	ProductID           uint    `gorm:"not null" json:"product_id" validate:"required"`
	Product             Product `gorm:"foreignKey:ProductID" json:"-"`
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
	ProductDetailsImage string  `json:"product_details_image"`
}
