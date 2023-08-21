//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/api"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/api/service"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/db"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/repository"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/usecase"
	"github.com/google/wire"
)

func InitializeServices(cfg *config.Config) (*api.ServiceServer, error) {
	wire.Build(
		//database connection
		db.ConnectDatabase,

		//service
		service.NewProductServiceServer,

		//database queries
		repository.NewProductRepository,

		//usecase
		usecase.NewProductUseCase,

		//server connection
		api.NewgrpcServer,
	)

	return &api.ServiceServer{}, nil
}
