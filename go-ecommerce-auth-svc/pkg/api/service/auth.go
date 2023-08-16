package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/pb"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase/interface"
)

type authServiceServer struct {
	authusecase services.AuthUseCase
	otpusecase  services.OTPUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase services.AuthUseCase, otpusecase services.OTPUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		authusecase: usecase,
		otpusecase:  otpusecase,
	}
}

func (cr *authServiceServer) UserSignup(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	newUserInfo := request.NewUserInfo{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}

	userDetails, err := cr.authusecase.UserSignUp(ctx, newUserInfo)
	if err != nil {
		return &pb.UserSignUpResponse{Status: http.StatusBadRequest}, errors.New("failed to create user")
		//c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to create user", err.Error(), nil))
	}

	//twilio otp send

	responseID, err := cr.otpusecase.TwilioSendOtp(ctx, "+91"+userDetails.Phone)
	if err != nil {
		return &pb.UserSignUpResponse{Status: http.StatusInternalServerError}, errors.New("failed to generate otp")

		// response := response.ErrorResponse(500, "failed to generate otp", err.Error(), nil)
		// c.JSON(http.StatusInternalServerError, response)
		// return

	}

	return &pb.UserSignUpResponse{Status: http.StatusOK, Responseid: responseID}, nil
	// response := response.SuccessResponse(200, "Success: Enter the otp and the response id", responseID)
	// c.JSON(http.StatusOK, response)

}

func (cr *authServiceServer) SignupOtpVerify(ctx context.Context, otpverify request.OTPVerify) (*pb.OtpVerifyResponse, error) {
	//var user domain.Users
	//var otpverify request.OTPVerify

	otpsession, err := cr.otpusecase.TwilioVerifyOTP(ctx, otpverify)
	if err != nil {
		return &pb.OtpVerifyResponse{Status: http.StatusBadRequest}, errors.New("Invalid Otp")

		// c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "Invalid Otp", err.Error(), nil))
		// return
	}

	// Call the OTPVerifyStatusManage method to update the verification status
	err = cr.authusecase.OTPVerifyStatusManage(ctx, otpsession)
	if err != nil {

		return &pb.OtpVerifyResponse{Status: http.StatusInternalServerError}, errors.New("Failed to update verification status")

		// response := response.ErrorResponse(500, "Failed to update verification status", err.Error(), nil)
		// c.JSON(http.StatusInternalServerError, response)
		// return
	}

	return &pb.OtpVerifyResponse{Status: http.StatusOK, Response: "user signin success"}, nil

	// response := response.SuccessResponse(200, "OTP validation Successfull..Account Created Successfully", nil)
	// c.JSON(200, response)
}
