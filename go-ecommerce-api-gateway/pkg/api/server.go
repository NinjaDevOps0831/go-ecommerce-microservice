package http

import (
	"github.com/ajujacob88/go-ecommerce-gin-clean-arch/pkg/api/handler"
	"github.com/ajujacob88/go-ecommerce-gin-clean-arch/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"honnef.co/go/tools/config"
)

type ServerHTTP struct {
	engine *gin.Engine
	Port   string
}

func NewServerHTTP(cfg *config.Config, userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler,
	productHandler *handler.ProductHandler,
	cartHandler *handler.CartHandler,
	orderHandler *handler.OrderHandler,
	paymentHandler *handler.PaymentHandler,
	couponHandler *handler.CouponHandler,

) *ServerHTTP {

	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	//add swagger - Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//setup routes
	routes.UserRoutes(engine.Group("/"), userHandler, productHandler, cartHandler, orderHandler, paymentHandler, couponHandler)
	routes.AdminRoutes(engine.Group("/admin"), adminHandler, productHandler, couponHandler, orderHandler)

	return &ServerHTTP{engine: engine, Port: cfg.Port}
}

func (sh *ServerHTTP) Start() {
	sh.engine.LoadHTMLGlob("views/*.html") //for loading the html page of razor pay
	sh.engine.Run(":3000")
}
