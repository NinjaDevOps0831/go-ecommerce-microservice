package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
)

type OTPUseCase interface {
	TwilioSendOtp(ctx context.Context, phoneNumber string) (string, error)
	TwilioVerifyOTP(ctx context.Context, otpverify request.OTPVerify) (domain.OTPSession, error)
}
