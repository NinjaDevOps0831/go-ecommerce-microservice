package domain

type Carts struct {
	ID              uint    `gorm:"primaryKey,uniqueIndex" json:"id"`
	UserID          uint    `gorm:"not null" json:"user_id" validate:"required"`
	Users           Users   `gorm:"foreignKey:UserID" json:"-"`
	SubTotal        float64 `gorm:"not null" json:"sub_total"`
	AppliedCouponID uint    `json:"applied_coupon_id"`
	DiscountAmount  float64 `json:"discount_amount"`
	TotalPrice      float64 `json:"total_price"`
}

type CartItems struct {
	ID               uint           `gorm:"primaryKey,uniqueIndex" json:"id" validate:"required"`
	CartID           uint           `gorm:"not null" json:"cart_id" validate:"required"`
	Carts            Carts          `gorm:"foreignKey:CartID" json:"-"`
	ProductDetailsID uint           `json:"product_details_id"`
	ProductDetails   ProductDetails `gorm:"foreignKey:ProductDetailsID" json:"-"`
	Quantity         uint           `json:"quantity"`
}
