package response

import "time"

//model structs used to output data

type AdminDataOutput struct {
	ID           uint
	UserName     string
	Email        string
	Phone        string
	IsSuperAdmin bool
}

type AdminDashboard struct {
	CompletedOrders   int     `json:"completed_orders"`
	PendingOrders     int     `json:"pending_orders"`
	CancelledOrders   int     `json:"cancelled_orders"`
	TotalOrders       int     `json:"total_orders"`
	TotalOrderedItems int     `json:"total_order_items"`
	OrderValue        float64 `json:"order_value"`
	CreditedAmount    float64 `json:"credited_amount"`
	PendingAmount     float64 `json:"pending_amount"`
	TotalUsers        int     `json:"total_users"`
	VerifiedUsers     int     `json:"verified_users"`
	OrderedUsers      int     `json:"ordered_users"`
}

type SalesReport struct {
	UserID                uint      `json:"user_id"`
	FirstName             string    `json:"first_name"`
	Email                 string    `json:"email"`
	OrderID               uint      `json:"order_id"`
	OrderTotalPrice       uint      `json:"order_total_price"`
	OrderDate             time.Time `json:"order_date"`
	OrderStatus           string    `json:"order_status"`
	AppliedCouponCode     string    `json:"applied_coupon_code"`
	AppliedCouponDiscount float64   `json:"applied_coupon_discount"`
	DeliveryStatus        string    `json:"delivery_status"`
	PaymentType           string    `json:"payment_type"`
}
