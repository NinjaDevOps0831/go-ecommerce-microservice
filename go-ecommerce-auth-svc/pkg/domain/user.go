package domain

import (
	"time"
)

type Users struct {
	//gorm.Model
	ID        uint   `json:"id" gorm:"primaryKey;unique"` //gorm.Model is used instead of id, created at, deleted at
	FirstName string `json:"first_name" gorm:"not null" binding:"required,min=3,max=18"`
	LastName  string `json:"last_name" binding:"required,max=15"`
	Email     string `json:"email" gorm:"unique,not null" binding:"required,email"`
	//Email    string `json:"email" gorm:"unique;not null"`
	Phone    string `json:"phone_no" gorm:"unique" binding:"required,min=10,max=10"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	//BlockStatus  bool      `json:"block_status" gorm:"not null;default:false"`
	//VerifyStatus bool      `json:"verify_status" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//UserInfo  UserInfo  `json:"user_info" gorm:"foreignKey:UsersID"`
}

type UserInfo struct {
	ID                uint `gorm:"primaryKey"`
	IsVerified        bool `json:"is_verified"`
	VerifiedAt        time.Time
	IsBlocked         bool `json:"is_blocked"`
	BlockedAt         time.Time
	BlockedBy         uint   `json:"blocked_by"`
	Admin             Admin  `gorm:"foreignKey:BlockedBy" json:"-"`
	ReasonForBlocking string `json:"reason_for_blocking"`
	UsersID           uint   `json:"users_id" json:"-"`
	Users             Users  `json:"users" gorm:"foreignKey:UsersID"`
	Check             string `json:"check"`
	//Users             Users  `gorm:"foreignKey:UsersID" json:"-"`

}

type UserAddress struct {
	ID          uint   `json:"id" gorm:"primaryKey;unique"`
	UserID      uint   `json:"user_id"`
	Users       Users  `gorm:"foreignKey:UserID" json:"-"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	District    string `json:"district"`
	State       string `json:"state"`
	Pincode     string `json:"pincode"`
	Landmark    string `json:"landmark"`
	//FullName    string `json:"full_name"`
	// Email       string `json:"email"`
	// Phone       string `json:"phone"`
}

/*

type Address struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Users     Users  `gorm:"foreignKey:UserID" json:"-"`
	HouseName string `json:"house_name"`
	Street    string `json:"street"`
	City      string `json:"city"`
	District  string `json:"district"`
	State     string `json:"state"`
	Pincode   string `json:"pincode"`
	Landmark  string `json:"landmark"`
}

*/
