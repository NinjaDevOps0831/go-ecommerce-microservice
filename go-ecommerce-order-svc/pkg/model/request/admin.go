package request

//model structs used for input

type NewAdminInfo struct {
	UserName string `json:"user_name" validate:"required"`
	//Email    string `json:"email" validate:"required,email"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required,min=10,max=10"`
	Password string `json:"password" validate:"required"`
}

type AdminLoginInfo struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
