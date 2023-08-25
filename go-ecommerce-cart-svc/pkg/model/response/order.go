package response

type UserOrder struct {
	AmountToPay           float64 `json:"amount_to_pay"`
	AppliedCouponID       uint
	AppliedCouponDiscount float64
}
