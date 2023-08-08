package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/pb"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/utils"
)

type AuthClient interface {
	UserSignup(context.Context, utils.BodySignUpuser) (*pb.RegisterResponse, error)
	UserSignupOtpVerify(context.Context, utils.Otpverify) (*pb.OtpVerifyResponse, error)
	UserLogin(context.Context, utils.BodyLogin) (*pb.LoginResponse, error)
	ShowAddress(ctx context.Context, id uint) (*pb.ShowUserAddressResponse, error)
	AddAddress(context.Context, utils.AddAddress, uint) (*pb.AddUserAddressResponse, error)

	//token validation
	AuthorizationMiddleware(string) (*pb.ValidateResponse, error)

	//Admin
	AdminSignup(context.Context, utils.BodySignUpuser) (*pb.OtpVerifyResponse, error)
	AdminLogin(context.Context, utils.BodyLogin) (*pb.LoginResponse, error)
}
