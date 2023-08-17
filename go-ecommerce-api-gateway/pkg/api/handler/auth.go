package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/handlerutil"

	client "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/client/interfaces"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Client client.AuthClient
}

func NewUserHandler(client client.AuthClient) AuthHandler {
	return AuthHandler{
		Client: client,
	}
}

// @title Ecommerce REST API Microservice
// @version 1.0
// @description Ecommerce REST API Microservice built using Go, PSQL, REST API following Clean Architecture.

// @contact
// name: Aju Jacob
// url: https://github.com/ajujacob88
// email: ajujacob88@gmail.com

// @license
// name: MIT
// url: https://opensource.org/licenses/MIT

//  If the @host is not specified in Swagger/OpenAPI documentation, it is assumed to be the same host where the API documentation is being served. In this case, if you access the Swagger UI from esmartstore.shop/swagger/index.html or www.esmartstore.shop/swagger/index.html, the API calls will be made to the same host by default.
// So, if you don't specify the @host, it will work fine as long as the Swagger UI and the API server are hosted on the same domain. The requests will be sent to the current host where the Swagger UI is being served from.
// read it  https://swagger.io/docs/specification/2-0/api-host-and-base-path/

// host localhost:3000    ----no need
// host esmartstore.shop     ----no need
// host www.esmartstore.shop   ----no need

// @Basepath /
// @Accept json
// @Produce json
// @Router / [get]

// UserSignup
// @Summary api for Signup a new user
// @ID Signup-user
// @Description Create a new user with the specified details.
// @Tags Users Signup
// @Accept json
// @Produce json
// @Param user_details body request.NewUserInfo true "New user Details"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /user/signup [post]
func (cr *AuthHandler) UserSignUp(c *gin.Context) {
	//var user domain.Users
	var newUserInfo request.NewUserInfo
	if err := c.BindJSON(&newUserInfo); err != nil {
		response := response.ErrorResponse(422, "unable to read the request body", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userDetails, err := cr.Client.UserSignup(context.Background(), newUserInfo)
	if err != nil {
		fmt.Println("debug check 1", userDetails)
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to create user", err.Error(), nil))
	}

	//twilio otp send

	// responseID, err := cr.otpUseCase.TwilioSendOtp(c.Request.Context(), "+91"+userDetails.Phone)
	// if err != nil {
	// 	response := response.ErrorResponse(500, "failed to generate otp", err.Error(), nil)

	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return

	// }
	response := response.SuccessResponse(200, "Success: Enter the otp and the response id", userDetails)
	c.JSON(http.StatusOK, response)

}

// SIGN UP OTP VERIFICATION
// SignupOtpVerify
// @Summary signup otp verification
// @ID Signup-otpverify-user
// @Description verify the otp of a user.
// @Tags Users otp verify
// @Accept json
// @Produce json
// @Param otpverify body request.OTPVerify true "OTP verification details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /user/signup/otp/verify [post]
func (cr *AuthHandler) SignupOtpVerify(c *gin.Context) {
	//var user domain.Users
	var otpverify request.OTPVerify
	if err := c.BindJSON(&otpverify); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to read the request body", err.Error(), nil))
		return
	}
	otpsession, err := cr.Client.UserSignupOtpVerify(context.Background(), otpverify)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "Invalid Otp", err.Error(), nil))
		return
	}

	// Call the OTPVerifyStatusManage method to update the verification status
	// err = cr.userUseCase.OTPVerifyStatusManage(c.Request.Context(), otpsession)
	// if err != nil {
	// 	response := response.ErrorResponse(500, "Failed to update verification status", err.Error(), nil)
	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	response := response.SuccessResponse(200, "OTP validation Successfull..Account Created Successfully", &otpsession)
	c.JSON(200, response)
}

// UserLogin
// @Summary User Login
// @ID user-login
// @Description user Login
// @Tags user
// @Accept json
// @Produce json
// @Param user_credentials body request.UserLoginEmail true "user Login Credentials"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/login/email [post]
func (cr *AuthHandler) UserLoginByEmail(c *gin.Context) {
	//receive data from request body
	var body request.UserLoginEmail

	if err := c.ShouldBindJSON(&body); err != nil {
		response := response.ErrorResponse(400, "Input is invalid", err.Error(), nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	//copy the body values to user
	// var user domain.Users
	// copier.Copy(&user, &body)

	res, err := cr.Client.UserLoginByEmail(context.Background(), body)
	if err != nil {
		response := response.ErrorResponse(400, "failed to login", err.Error(), nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// generate token using jwt in map
	// tokenString, err := auth.GenerateJWT(res.Token)
	// if err != nil {
	// 	response := response.ErrorResponse(500, "faild to generate jwt", err.Error(), nil)
	// 	c.JSON(http.StatusInternalServerError, response)
	// 	return
	// }

	c.SetCookie("UserAuth", res.Token, 60*60, "", "", false, true)

	//response := response.SuccessResponse(200, "successfully logged in", tokenString["accessToken"])
	response := response.SuccessResponse(200, "successfully logged in", res.User)
	c.JSON(http.StatusOK, response)
}

// AddAddress
// @Summary User can add the user address
// @ID add-address
// @Description Add address
// @Tags user
// @Accept json
// @Produce json
// @Param user_address body request.UserAddressInput true "User address"
// @Success 201 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /user/addresses/ [post]
func (cr *AuthHandler) AddAddress(c *gin.Context) {
	var userAddressInput request.UserAddressInput
	if err := c.Bind(&userAddressInput); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.ErrorResponse(422, "unable to read the request body", err.Error(), nil))
		return
	}

	// //both c.Get and c.Value can be used to retrieve data from context.. here i just used c.Value and i have also used the function GetUserIdFromContext
	// id, valuebool := c.Get("userID")
	// userid, err := strconv.Atoi(fmt.Sprintf("%v", id))

	userID, err := handlerutil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "unable to fetch user id from context", err.Error(), nil))
		return
	}

	address, err := cr.Client.AddAddress(context.Background(), userAddressInput, uint(userID))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to add the address", err.Error(), nil))
		return
	}

	c.JSON(http.StatusCreated, response.SuccessResponse(201, "Succesfully added the address", &address))

}

// Create Admin - SuperAdmin can create a new admin from admin panel
// @Summary Create a new admin from admin panel
// @ID create-admin
// @Description Super admin can create a new admin from admin panel
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_details body request.NewAdminInfo true "New Admin Details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 422 {object} response.Response
// @Router /admin/admins [post]
func (cr *AuthHandler) CreateAdmin(c *gin.Context) {
	var newAdminInfo request.NewAdminInfo
	fmt.Println("debug check 1 newadmin info is", newAdminInfo)
	if err := c.Bind(&newAdminInfo); err != nil {
		//The 422 status code is often used in API scenarios where clients submit data that fails validation, such as missing required fields, invalid data formats, or conflicting information.
		response := response.ErrorResponse(422, "unable to read the request body", err.Error(), nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//finding out the admin id of the admin who is trying to create the new user., if the admin is super admin, then only he can able to create a new admin.
	adminID, err := handlerutil.GetAdminIdFromContext(c)
	fmt.Println("Admin ID is(for superuser check)", adminID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to fetch the admin ID", err.Error(), nil))
		return
	}
	//Now call the create admin method from admin usecase. The admin data will be saved to domain.admin after the succesful execution of the function
	newAdminOutput, err := cr.Client.CreateAdmin(c.Request.Context(), newAdminInfo, uint32(adminID))

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to create the admin", err.Error(), nil))
		return
	}

	//if no error, then  201 status as new admin is created succesfully
	c.JSON(http.StatusCreated, response.SuccessResponse(201, "admin created successfully", newAdminOutput))

}

// AdminLogin
// @Summary Admin Login for the admin
// @ID admin-login
// @Description Admin Login forthe admin review
// @Tags Admin
// @Accept json
// @Produce json
// @Param admin_credentials body request.AdminLoginInfo true "Admin Login Credentials"
// @Success 200 {object} response.Response
// @Failure 422 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /admin/login [post]
func (cr *AuthHandler) AdminLogin(c *gin.Context) {
	//receive the data from request body
	var body request.AdminLoginInfo
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.Response{StatusCode: 422, Message: "unable to process the request", Errors: err.Error(), Data: nil})
		return
	}
	fmt.Println("debug test - body", body)

	res, err := cr.Client.AdminLogin(context.Background(), body)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to login", err.Error(), nil))
		return
	}
	c.SetSameSite(http.SameSiteLaxMode) //sets the SameSite attribute of the cookie to "Lax" mode. It is a security measure that helps protect against certain types of cross-site request forgery (CSRF) attacks.
	c.SetCookie("AdminAuth", res.Token, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, response.SuccessResponse(200, "Succesfully Logged in", res.AdminDataOutput))
}
