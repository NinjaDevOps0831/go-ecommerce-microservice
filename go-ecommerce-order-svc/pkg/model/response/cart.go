package response

//model structs used for input

type CartItems struct {
	ProductItemID    uint
	Brand            string
	Name             string
	Model            string
	Quantity         uint
	ProductItemImage string
	Price            float64
	Total            float64
}

type ViewCart struct {
	CartItemsAll []CartItems
	// SubTotal        float64
	// AppliedCouponID uint
	// DiscountAmount  float64
	// TotalPrice      float64
	CartDetails // used struct embedding,instead of creating the above fields again,

}

// this is for view cart from cart repo viewcart function
type CartDetails struct {
	ID              int `json:"-"`
	SubTotal        float64
	AppliedCouponID uint
	DiscountAmount  float64
	TotalPrice      float64
}
