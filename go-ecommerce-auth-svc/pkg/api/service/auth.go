package service

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/pb"
	usecase "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/useCase"
)

type authServiceServer struct {
	usecase usecase.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase usecase.AuthUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		usecase: usecase,
	}
}
