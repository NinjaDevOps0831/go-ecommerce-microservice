package interfaces

import (
	"context"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
)

type OTPRepository interface {
	SaveOTP(ctx context.Context, resp string, phoneNumber string) error
	RetrieveOtpSession(ctx context.Context, otpverify request.OTPVerify) (domain.OTPSession, error)
}
