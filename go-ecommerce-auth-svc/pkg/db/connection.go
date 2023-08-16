package db

import (
	"fmt"
	"log"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
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

		//user tables
		&domain.Users{},
		&domain.UserInfo{},
		&domain.UserAddress{},

		//otp table for temporary storage purpose
		&domain.OTPSession{},

		//admin tables
		&domain.Admin{}, //By default, GORM automatically pluralizes the table name based on the struct name. That's why the Admin struct becomes the "admins" table.

		// //product tables
		// &domain.ProductCategory{},
		// &domain.Product{},
		// &domain.ProductBrand{},
		// &domain.ProductDetails{},

		// //cart tables
		// &domain.Carts{},
		// &domain.CartItems{},

		// //order tables
		// &domain.Order{},
		// &domain.OrderLine{},
		// &domain.OrderStatus{},
		// &domain.DeliveryStatus{},
		// &domain.OrderReturn{},

		// //payment tables
		// &domain.PaymentMethodInfo{},
		// &domain.PaymentDetails{},
		// &domain.PaymentStatus{},

		// //coupon table
		// &domain.Coupon{},
		// domain.CouponUsed{},
	)

	if err != nil {
		return nil, err
	}

	return db, dbErr
}
