package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/response"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
)

type AuthRepository interface {
	UserSignup(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error)
	//UserSignupOtpVerify(ctx context.Context, otpverify request.OTPVerify) (*pb.OtpVerifyResponse, error)
	//UserLoginByEmail(ctx context.Context, body request.UserLoginEmail) (*pb.LoginResponse, error)

	FindByEmail(ctx context.Context, Email string) (domain.Users, error)

	BlockStatus(ctx context.Context, userId uint) (bool, error)

	//ShowAddress(ctx context.Context, id uint) (*pb.ShowUserAddressResponse, error)
	AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (domain.UserAddress, error)

	//token validation
	//AuthorizationMiddleware(string) (*pb.ValidateResponse, error)

	//Admin
	CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo, adminID uint32) (domain.Admin, error)
	//AdminLogin(ctx context.Context, body request.AdminLoginInfo) (*pb.AdminLoginResponse, error)
	FindAdmin(ctx context.Context, email string) (domain.Admin, error)
}
