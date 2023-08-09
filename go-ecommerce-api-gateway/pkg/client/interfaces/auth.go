package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/pb"
)

type AuthClient interface {
	UserSignup(ctx context.Context, newUserInfo request.NewUserInfo) (*pb.UserSignUpResponse, error)
	UserSignupOtpVerify(ctx context.Context, otpverify request.OTPVerify) (*pb.OtpVerifyResponse, error)
	UserLoginByEmail(ctx context.Context, body request.UserLoginEmail) (*pb.LoginResponse, error)
	//ShowAddress(ctx context.Context, id uint) (*pb.ShowUserAddressResponse, error)
	AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (*pb.AddUserAddressResponse, error)

	//token validation
	//AuthorizationMiddleware(string) (*pb.ValidateResponse, error)

	//Admin
	CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo) (*pb.AdminSignupResponse, error)
	AdminLogin(ctx context.Context, body request.AdminLoginInfo) (*pb.AdminLoginResponse, error)
}
