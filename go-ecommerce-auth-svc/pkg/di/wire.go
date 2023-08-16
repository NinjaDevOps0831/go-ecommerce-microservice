//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/api"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/api/service"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/db"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/repository"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase"
	"github.com/google/wire"
)

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {
	wire.Build(
		//database connection
		db.ConnectDatabase,

		//service
		service.NewAuthServiceServer,

		//database queries
		repository.NewAuthRepository,
		repository.NewOTPRepository,

		//usecase
		usecase.NewUserUseCase,
		usecase.NewOTPUseCase,

		//server connection
		api.NewgrpcServer,
	)

	return &api.ServiceServer{}, nil
}
