package db

import (
	"fmt"
	"log"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-product-svc/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if dbErr != nil {
		log.Fatalln(dbErr)
	}

	err := db.AutoMigrate(

		//product tables
		&domain.ProductCategory{},
		&domain.Product{},
		&domain.ProductBrand{},
		&domain.ProductDetails{},
	)

	if err != nil {
		return nil, err
	}

	return db, dbErr
}
