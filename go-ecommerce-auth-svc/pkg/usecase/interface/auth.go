package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
)

type AuthUseCase interface {
	UserSignUp(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error)

	LoginWithEmail(ctx context.Context, user request.UserLoginEmail) (domain.Users, error)

	//	FindByEmail(ctx context.Context, Email string) (domain.Users, error)

	AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID int) (domain.UserAddress, error)

	CreateAdmin(ctx context.Context, newAdmin request.NewAdminInfo, adminID int) (domain.Admin, error)

	AdminLogin(ctx context.Context, input request.AdminLoginInfo) (string, response.AdminDataOutput, error)
}
