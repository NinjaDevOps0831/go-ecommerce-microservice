package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID                    uint              `json:"id" gorm:"primaryKey"`
	UserID                uint              `json:"user_id"`
	Users                 Users             `gorm:"foreignKey: UserID" json:"-"`
	OrderDate             time.Time         `json:"order_date"`
	PaymentMethodInfoID   uint              `json:"payment_method_info_id"`
	PaymentMethodInfo     PaymentMethodInfo `gorm:"foreignKey:PaymentMethodInfoID" json:"-"`
	ShippingAddressID     uint              `json:"shipping_address_id"`
	UserAddress           UserAddress       `gorm:"foreignKey: ShippingAddressID" json:"-"`
	OrderTotalPrice       float64           `json:"order_total_price"`
	OrderStatusID         uint              `json:"order_status_id"`
	OrderStatus           OrderStatus       `gorm:"foreignKey: OrderStatusID" json:"-"`
	AppliedCouponID       uint              `json:"applied_coupon_id"`
	AppliedCouponDiscount float64           `json:"applied_coupon_discount"`
	DeliveryStatusID      uint              `json:"delivery_status_id"`
	DeliveryStatus        DeliveryStatus    `gorm:"foreignKey: DeliveryStatusID" json:"-"`
	DeliveredAt           *time.Time        `json:"delivered_at"` //By using a nullable time type like *time.Time, the used_at field will be correctly mapped as null when it is not present in the database, rather than being assigned a default value. It can be useful in scenarios where the field can have a nullable value.
}

type OrderStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type DeliveryStatus struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
}

type OrderLine struct {
	ID               uint           `gorm:"primaryKey"`
	ProductDetailsID uint           `json:"product_details_id"`
	ProductDetails   ProductDetails `gorm:"ForeignKey: ProductDetailsID" json:"-"`
	OrderID          uint           `json:"order_id"`
	Order            Order          `gorm:"ForeignKey: OrderID" json:"-"`
	Quantity         int            `json:"quantity"`
	Price            float64        `json:"price"`
}

type OrderReturn struct {
	gorm.Model
	OrderID      uint    `json:"order_id"`
	Order        Order   `gorm:"ForeignKey: OrderID" json:"-"`
	ReturnReason string  `json:"return_reason" gorm:"not null"`
	RefundAmount float64 `json:"refund_amount" gorm:"not null"`

	IsApproved   bool      `json:"is_approved"`
	ApprovalDate time.Time `json:"approval_date"`
	ReturnDate   time.Time `json:"return_date"`
	AdminComment string    `json:"admin_comment"`
}
