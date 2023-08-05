package main

import (
	"log"

	_ "github.com/ajujacob88/go-ecommerce-gin-clean-arch/cmd/api/docs" //this is for swagger docs, swagger will work only if this is here, also give _ before the code, otherwise this will gone when saved(because it doesnt used directly)
	"github.com/ajujacob88/go-ecommerce-gin-clean-arch/pkg/config"
	"github.com/ajujacob88/go-ecommerce-gin-clean-arch/pkg/di"
)

func main() {
	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config: ", configErr)
	}

	server, diErr := di.InitializeAPI(config)
	if diErr != nil {
		log.Fatal("cannot start server: ", diErr)
	} else {
		server.Start()
	}
}
