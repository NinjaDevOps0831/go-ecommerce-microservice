package request

type PlaceOrder struct {
	PaymentMethodID int `json:"payment_method_id" binding:"required"`
	AddressID       int `json:"address_id" binding:"required"`
}

type ReturnRequest struct {
	OrderID      uint   `json:"order_id"`
	ReturnReason string `json:"resturn_reason"`
}

type UpdateOrderStatuses struct {
	OrderID          uint `json:"order_id"`
	OrderStatusID    uint `json:"order_status_id"`
	DeliveryStatusID uint `json:"delivery_status_id"`
}
