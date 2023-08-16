package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
)

type AuthRepository interface {
	FindUser(ctx context.Context, newUser request.NewUserInfo) (domain.Users, error)

	UserSignUp(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error)
	//UserSignupOtpVerify(ctx context.Context, otpverify request.OTPVerify) (*pb.OtpVerifyResponse, error)
	//UserLoginByEmail(ctx context.Context, body request.UserLoginEmail) (*pb.LoginResponse, error)

	OTPVerifyStatusManage(ctx context.Context, otpsession domain.OTPSession) error

	FindByEmail(ctx context.Context, Email string) (domain.Users, error)

	BlockStatus(ctx context.Context, userId uint) (bool, error)

	//ShowAddress(ctx context.Context, id uint) (*pb.ShowUserAddressResponse, error)
	AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (domain.UserAddress, error)

	//token validation
	//AuthorizationMiddleware(string) (*pb.ValidateResponse, error)

	//Admin

	IsSuperAdmin(ctx context.Context, adminID uint) (bool, error)

	CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo) (domain.Admin, error)
	//AdminLogin(ctx context.Context, body request.AdminLoginInfo) (*pb.AdminLoginResponse, error)
	FindAdmin(ctx context.Context, email string) (domain.Admin, error)
}
