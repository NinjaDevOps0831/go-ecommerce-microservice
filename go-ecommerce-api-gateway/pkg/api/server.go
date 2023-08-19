package api

import (
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/handler"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/routes"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	Engine *gin.Engine
	Port   string
}

func NewServerHTTP(cfg *config.Config, authHandler handler.AuthHandler, productHandler handler.ProductHandler,

) (*ServerHTTP, error) {

	engine := gin.New()

	// Use logger from Gin
	engine.Use(gin.Logger())

	//add swagger - Swagger docs
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//setup routes
	routes.UserRoutes(engine.Group("/"), authHandler, productHandler)
	routes.AdminRoutes(engine.Group("/admin"), authHandler, productHandler)

	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})

	return &ServerHTTP{Engine: engine, Port: cfg.Port}, nil
}

func (sh *ServerHTTP) Start() {
	//sh.Engine.LoadHTMLGlob("views/*.html") //for loading the html page of razor pay
	sh.Engine.Run(sh.Port)
}
