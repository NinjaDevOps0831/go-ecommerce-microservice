package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/auth"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/pb"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase/interface"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/copier"
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

func (cr *authServiceServer) UserLoginByEmail(ctx context.Context, body request.UserLoginEmail) (*pb.LoginResponse, error) {
	//receive data from request body
	//var body request.UserLoginEmail

	//copy the body values to user
	var user domain.Users
	copier.Copy(&user, &body)

	// get user from database and check password in usecase
	user, err := cr.authusecase.LoginWithEmail(ctx, body)
	if err != nil {
		return &pb.LoginResponse{Status: http.StatusBadRequest}, errors.New("failed to login")

	}

	// generate token using jwt in map
	tokenString, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return &pb.LoginResponse{Status: http.StatusInternalServerError}, errors.New("faild to generate jwt")

	}
	var c *gin.Context

	c.SetCookie("UserAuth", tokenString, 60*60, "", "", false, true)

	data := &pb.User{
		FirstName: user.FirstName,
	}
	return &pb.LoginResponse{Status: http.StatusOK, User: data,
		Token: tokenString}, nil

	// response := response.SuccessResponse(200, "successfully logged in", user.FirstName)
	// c.JSON(http.StatusOK, response)
}

func (cr *authServiceServer) AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (*pb.AddUserAddressResponse, error) {
	//var userAddressInput request.UserAddressInput
	// if err := c.Bind(&userAddressInput); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to read the request body", err.Error(), nil))
	// 	return
	// }

	// userID, err := handlerutil.GetUserIdFromContext(ctx)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "unable to fetch user id from context", err.Error(), nil))
	// 	return
	// }

	address, err := cr.authusecase.AddAddress(ctx, userAddressInput, userID)
	if err != nil {
		return &pb.AddUserAddressResponse{Status: http.StatusBadRequest}, errors.New("failed to add the address")

	}

	data := &pb.UserAddressOutput{
		Id:          uint32(address.ID),
		Userid:      uint32(address.UserID),
		HouseNumber: address.HouseNumber,
		Street:      address.Street,
		City:        address.City,
		District:    address.District,
		State:       address.State,
		Pincode:     address.Pincode,
	}
	return &pb.AddUserAddressResponse{Status: http.StatusCreated, UserAddressOutput: data}, errors.New("Succesfully added the address")

}

func (cr *authServiceServer) CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo, adminID uint32) (*pb.AdminSignupResponse, error) {
	//var newAdminInfo request.NewAdminInfo

	//finding out the admin id of the admin who is trying to create the new user., if the admin is super admin, then only he can able to create a new admin.
	//adminID, err := handlerutil.GetAdminIdFromContext(c)
	fmt.Println("Admin ID is(for superuser check)", adminID)

	//Now call the create admin method from admin usecase. The admin data will be saved to domain.admin after the succesful execution of the function
	newAdminOutput, err := cr.authusecase.CreateAdmin(ctx, newAdminInfo, uint(adminID))

	if err != nil {
		return &pb.AdminSignupResponse{Status: http.StatusBadRequest}, errors.New("failed to create the admin")

	}

	data := &pb.Admin{
		Id:           uint32(newAdminOutput.ID),
		UserName:     newAdminOutput.UserName,
		Email:        newAdminOutput.Email,
		PhoneNo:      newAdminOutput.Phone,
		IsSuperAdmin: newAdminOutput.IsBlocked,
		IsBlocked:    newAdminOutput.IsBlocked,
	}
	return &pb.AdminSignupResponse{Status: http.StatusCreated, Admin: data}, nil

}

func (cr *authServiceServer) AdminLogin(ctx context.Context, body request.AdminLoginInfo) (*pb.AdminLoginResponse, error) {
	//receive the data from request body
	//var body request.AdminLoginInfo

	//call the adminlogin method of the adminusecase to login as an admin
	tokenString, adminDataInModel, err := cr.authusecase.AdminLogin(ctx, body)
	if err != nil {
		return &pb.AdminLoginResponse{Status: http.StatusBadRequest}, errors.New("failed to login")

	}
	var c *gin.Context
	c.SetSameSite(http.SameSiteLaxMode) //sets the SameSite attribute of the cookie to "Lax" mode. It is a security measure that helps protect against certain types of cross-site request forgery (CSRF) attacks.
	c.SetCookie("AdminAuth", tokenString, 3600*24*30, "", "", false, true)

	data := &pb.AdminDataOutput{
		Id:           uint32(adminDataInModel.ID),
		UserName:     adminDataInModel.UserName,
		Email:        adminDataInModel.Email,
		Phone:        adminDataInModel.Phone,
		IsSuperAdmin: adminDataInModel.IsSuperAdmin,
	}
	return &pb.AdminLoginResponse{Status: http.StatusOK, AdminDataOutput: data, Token: tokenString}, nil

}
