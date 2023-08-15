package repository

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type authDatabse struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{
		DB: DB}
}
