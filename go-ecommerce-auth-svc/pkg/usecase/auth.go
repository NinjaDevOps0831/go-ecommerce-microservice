package usecase

import (
	interfaces "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/repository/interface"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase/interface"
)

type AuthUseCase struct {
	authRepo interfaces.AuthRepository
}

func NewUserUseCase(repo interfaces.AuthRepository) services.AuthUseCase {
	return &authUseCase{
		authRepo: repo,
	}

}
