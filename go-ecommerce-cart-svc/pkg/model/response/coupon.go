package response

import "time"

type ViewCoupons struct {
	CouponName        string     `json:"coupon_name"`
	CouponCode        string     `json:"coupon_code"`
	MinOrderValue     float64    `json:"min_order_value"`
	DiscountPercent   float64    `json:"discount_percent"`
	DiscountMaxAmount float64    `json:"discount_max_amount"`
	ValidTill         time.Time  `json:"valid_till"`
	Description       string     `json:"description"`
	UsedAt            *time.Time `json:"used_at"` //By using a nullable time type like *time.Time, the used_at field will be correctly mapped as null when it is not present in the database, rather than being assigned a default value. It can be useful in scenarios where the field can have a nullable value.
}
