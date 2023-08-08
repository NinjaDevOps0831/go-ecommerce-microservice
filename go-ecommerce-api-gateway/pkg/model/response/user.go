package response

//model structs output

type UserDataOutput struct {
	ID        uint   `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type ShowAddress struct {
	// FullName    string `json:"full_name"`
	// Email       string `json:"email"`
	// Phone       string `json:"phone"`
	ID          uint   `json:"address_id"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	State       string `json:"state"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}
