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

func (cr *authServiceServer) UserSignUp(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	fmt.Println("debug test -3 - req is", req)
	newUserInfo := request.NewUserInfo{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}
	fmt.Println("debug test -4 - newuserinfo is", newUserInfo)

	userDetails, err := cr.authusecase.UserSignUp(ctx, newUserInfo)
	if err != nil {
		fmt.Println("debug test -5 - userdetails is", userDetails, "error is", err)
		return &pb.UserSignUpResponse{Status: http.StatusBadRequest}, errors.New("failed to create user")
		//c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to create user", err.Error(), nil))
	}

	fmt.Println("debug test -6 - userdetails is", userDetails)

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

func (cr *authServiceServer) SignupOtpVerify(ctx context.Context, req *pb.OtpVerifyRequest) (*pb.OtpVerifyResponse, error) {
	//var user domain.Users
	//var otpverify request.OTPVerify
	otpverify := request.OTPVerify{
		OTP:   req.GetOtp(),
		OtpId: req.GetOtpId(),
	}

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

func (cr *authServiceServer) UserLoginByEmail(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	//receive data from request body
	//var body request.UserLoginEmail

	body := request.UserLoginEmail{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

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

func (cr *authServiceServer) AddAddress(ctx context.Context, req *pb.AddUserAddressRequest) (*pb.AddUserAddressResponse, error) {
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

	userAddressInput := request.UserAddressInput{
		HouseNumber: req.GetHouseNumber(),
		Street:      req.GetStreet(),
		City:        req.GetCity(),
		District:    req.GetDistrict(),
		State:       req.GetState(),
		Pincode:     req.GetPincode(),
		Landmark:    req.GetLandmark(),
	}

	userID := uint(req.GetUserid())

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

func (cr *authServiceServer) CreateAdmin(ctx context.Context, req *pb.AdminSignupRequest) (*pb.AdminSignupResponse, error) {
	//var newAdminInfo request.NewAdminInfo

	//finding out the admin id of the admin who is trying to create the new user., if the admin is super admin, then only he can able to create a new admin.
	//adminID, err := handlerutil.GetAdminIdFromContext(c)
	//fmt.Println("Admin ID is(for superuser check)", adminID)

	newAdminInfo := request.NewAdminInfo{
		UserName: req.GetUserName(),
		Email:    req.GetEmail(),
		Phone:    req.GetPhone(),
		Password: req.GetPassword(),
	}

	adminID := req.GetAdminID()

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

func (cr *authServiceServer) AdminLogin(ctx context.Context, req *pb.LoginRequest) (*pb.AdminLoginResponse, error) {
	//receive the data from request body
	//var body request.AdminLoginInfo
	body := request.AdminLoginInfo{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	fmt.Println("debug test -3 - body is", body)

	//call the adminlogin method of the adminusecase to login as an admin
	tokenString, adminDataInModel, err := cr.authusecase.AdminLogin(ctx, body)
	if err != nil {
		fmt.Println("debug test -4 - err is", err)
		return &pb.AdminLoginResponse{Status: http.StatusBadRequest}, errors.New("failed to login")

	}
	fmt.Println("debug test -5 - tokenstring is", tokenString)
	// var c *gin.Context
	// c.SetSameSite(http.SameSiteLaxMode) //sets the SameSite attribute of the cookie to "Lax" mode. It is a security measure that helps protect against certain types of cross-site request forgery (CSRF) attacks.
	// c.SetCookie("AdminAuth", tokenString, 3600*24*30, "", "", false, true)

	data := &pb.AdminDataOutput{
		Id:           uint32(adminDataInModel.ID),
		UserName:     adminDataInModel.UserName,
		Email:        adminDataInModel.Email,
		Phone:        adminDataInModel.Phone,
		IsSuperAdmin: adminDataInModel.IsSuperAdmin,
	}
	return &pb.AdminLoginResponse{Status: http.StatusOK, AdminDataOutput: data, Token: tokenString}, nil

}
