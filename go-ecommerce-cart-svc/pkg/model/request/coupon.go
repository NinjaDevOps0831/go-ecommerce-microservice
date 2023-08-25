package request

import "time"

type ApplyCoupon struct {
	CouponCode string `json:"coupon_code" binding:"required"`
}

type Coupon struct {
	CouponName        string    `json:"coupon_name"`
	CouponCode        string    `json:"-"` //coupon code is generated, so no need to show it in swagger
	MinOrderValue     float64   `json:"min_order_value"`
	DiscountPercent   float64   `json:"discount_percent"`
	DiscountMaxAmount float64   `json:"discount_max_amount"`
	ValidTill         time.Time `json:"valid_till"` //Here's an example of how you can provide the ValidTill value in the Indian time zone: "2023-06-30T23:30:00+05:30"
	Description       string    `json:"description"`
}
