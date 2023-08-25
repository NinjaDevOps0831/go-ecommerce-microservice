package domain

import "time"

type Admin struct {
	ID       uint   `json:"id" gorm:"primaryKey,index"`
	UserName string `json:"user_name" gorm:"uniqueIndex"`
	//Email        string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Email        string `json:"email" gorm:"unique,not null" binding:"required,email"`
	Phone        string `json:"phone_no" gorm:"uniqueIndex"`
	Password     string `json:"password" gorm:"not null"`
	IsSuperAdmin bool   `json:"is_super_admin"`
	IsBlocked    bool   `josn:"is_blocked"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
