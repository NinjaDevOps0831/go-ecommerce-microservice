package request

type PaymentVerification struct {
	UserID          int
	OrderID         int
	RazorpayOrderID string
	PaymentRef      string
	Total           float64
}
