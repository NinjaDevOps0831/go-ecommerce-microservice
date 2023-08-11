//go:build wireinject
// +build wireinject

package di

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/handler"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/client"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/config"

	"github.com/google/wire"
)

//this provide config func was written as suggested by chatgpt while encountered an error in running wire

func ProvideConfig() (*config.Config, error) {
	return config.LoadConfig()
}

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(

		ProvideConfig, //this provide config func was written as suggested by chatgpt while encountered an error in running wire

		client.NewAuthClient,

		handler.NewUserHandler,

		//server connection
		api.NewServerHTTP,
	)

	return &api.ServerHTTP{}, nil
}
