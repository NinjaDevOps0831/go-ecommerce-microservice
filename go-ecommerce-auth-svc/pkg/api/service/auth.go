package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/pb"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase/interface"
)

type authServiceServer struct {
	usecase services.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase services.AuthUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		usecase: usecase,
	}
}

func (c *authServiceServer) UserSignup(ctx context.Context, req *pb.UserSignUpRequest) (*pb.UserSignUpResponse, error) {
	newUserInfo := request.NewUserInfo{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}

	userDetails, err := c.usecase.UserSignUp(ctx, newUserInfo)
	if err != nil {
		return &pb.UserSignUpResponse{Status: http.StatusBadRequest}, errors.New("failed to create user")
		//c.JSON(http.StatusBadRequest, response.ErrorResponse(400, "failed to create user", err.Error(), nil))
	}

	//twilio otp send

	responseID, err := c.usecase.TwilioSendOtp(c.Request.Context(), "+91"+userDetails.Phone)
	if err != nil {
		response := response.ErrorResponse(500, "failed to generate otp", err.Error(), nil)

		c.JSON(http.StatusInternalServerError, response)
		return

	}
	response := response.SuccessResponse(200, "Success: Enter the otp and the response id", responseID)
	c.JSON(http.StatusOK, response)

}
