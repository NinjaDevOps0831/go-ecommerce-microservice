package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
)

type AuthUseCase interface {
	UserSignUp(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error)

	OTPVerifyStatusManage(ctx context.Context, otpsession domain.OTPSession) error

	LoginWithEmail(ctx context.Context, user request.UserLoginEmail) (domain.Users, error)

	//	FindByEmail(ctx context.Context, Email string) (domain.Users, error)

	AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (domain.UserAddress, error)

	CreateAdmin(ctx context.Context, newAdmin request.NewAdminInfo, adminID uint) (domain.Admin, error)

	AdminLogin(ctx context.Context, input request.AdminLoginInfo) (string, response.AdminDataOutput, error)

	//TwilioSendOtp(ctx context.Context, phoneNumber string) (string, error)
	//TwilioVerifyOTP(ctx context.Context, otpverify request.OTPVerify) (domain.OTPSession, error)
}
