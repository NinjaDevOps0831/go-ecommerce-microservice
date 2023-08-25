package request

//model structs input

type BlockUser struct {
	UserID int    `json:"user_id"`
	Reason string `json:"reason"`
}

type NewUserInfo struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required,min=10,max=10"`
	//Email    string `json:"email" `
	//Phone    string `json:"phone" `
	Password string `json:"password" validate:"required"`
}

type UserLoginEmail struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserAddressInput struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	State       string `json:"state"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
}

type UserCredentials struct {
	Email    string `json:"email"`
	PhoneNum string `json:"phonenum"`
}
