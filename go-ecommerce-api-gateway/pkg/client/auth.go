package client

import (
	"context"
	"fmt"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/client/interfaces"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authClient struct {
	client pb.AuthServiceClient
}

func NewAuthClient(cfg *config.Config) (interfaces.AuthClient, error) {

	gcc, err := grpc.Dial(cfg.AuthServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewAuthServiceClient(gcc)

	return &authClient{
		client: client,
	}, nil
}

func (cr *authClient) UserSignup(ctx context.Context, newUserInfo request.NewUserInfo) (*pb.UserSignUpResponse, error) {
	fmt.Println("debug test 2 - newuserinfo", newUserInfo)
	res, err := cr.client.UserSignUp(ctx, &pb.UserSignUpRequest{
		FirstName: newUserInfo.FirstName,
		LastName:  newUserInfo.LastName,
		Email:     newUserInfo.Email,
		Phone:     newUserInfo.Phone,
		Password:  newUserInfo.Password,
	})
	fmt.Println("debug test 3 - res", res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) UserSignupOtpVerify(ctx context.Context, otpverify request.OTPVerify) (*pb.OtpVerifyResponse, error) {
	res, err := cr.client.SignupOtpVerify(ctx, &pb.OtpVerifyRequest{
		Otp:   otpverify.OTP,
		OtpId: otpverify.OtpId,
	})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) UserLoginByEmail(ctx context.Context, body request.UserLoginEmail) (*pb.LoginResponse, error) {
	res, err := cr.client.UserLoginByEmail(ctx, &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (*pb.AddUserAddressResponse, error) {
	res, err := cr.client.AddAddress(ctx, &pb.AddUserAddressRequest{
		Userid:      uint32(userID),
		HouseNumber: userAddressInput.HouseNumber,
		Street:      userAddressInput.Street,
		City:        userAddressInput.City,
		District:    userAddressInput.District,
		State:       userAddressInput.State,
		Pincode:     userAddressInput.Pincode,
		Landmark:    userAddressInput.Landmark,
	})

	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo, adminID uint32) (*pb.AdminSignupResponse, error) {
	res, err := cr.client.CreateAdmin(ctx, &pb.AdminSignupRequest{
		UserName: newAdminInfo.UserName,
		Email:    newAdminInfo.Email,
		Phone:    newAdminInfo.Phone,
		Password: newAdminInfo.Password,
		AdminID:  adminID,
	})

	if err != nil {
		return res, err
	}
	return res, nil
}

func (cr *authClient) AdminLogin(ctx context.Context, body request.AdminLoginInfo) (*pb.AdminLoginResponse, error) {
	res, err := cr.client.AdminLogin(ctx, &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})
	fmt.Println("debug test 2 - admin login creds", res)

	if err != nil {
		return res, err
	}
	return res, nil
}
