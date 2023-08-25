package domain

import "time"

type PaymentMethodInfo struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	PaymentType    string `json:"payment_type" gorm:"not null"`
	BlockStatus    bool   `json:"block_status" gorm:"not null;default:false"`
	MaxAmountLimit uint   `json:"max_amount_limit" gorm:"not null"`
}

type PaymentDetails struct {
	ID                  uint              `json:"id" gorm:"primaryKey"`
	OrderID             uint              `json:"order_id"`
	Order               Order             `gorm:"foreignKey: OrderID" json:"-"`
	OrderTotalPrice     float64           `json:"order_total_price"`
	PaymentMethodInfoID uint              `json:"payment_method_info_id"`
	PaymentMethodInfo   PaymentMethodInfo `gorm:"foreignKey: PaymentMethodInfoID" json:"-"`
	PaymentStatusID     uint              `json:"payment_status_id"`
	PaymentStatus       PaymentStatus     `gorm:"foreignKey:PaymentStatusID" json:"-"`
	PaymentRef          string            `json:"payment_ref"`
	UpdatedAt           time.Time         `json:"updated_at"`
}

type PaymentStatus struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	PaymentStatus string `json:"payment_status"`
}
