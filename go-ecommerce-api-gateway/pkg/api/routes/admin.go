package routes

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/handler"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(
	api *gin.RouterGroup,
	authHandler handler.AuthHandler,
	// productHandler handler.ProductHandler,
	// couponHandler handler.CouponHandler,
	// orderHandler handler.OrderHandler,

) {

	api.POST("/login", authHandler.AdminLogin)

	api.Use(middleware.AdminAuth) //Middleware functions in Gin are handlers that are executed before reaching the final handler for a particular route. They provide a way to perform common tasks, such as authentication, logging, or data preprocessing, for multiple routes or groups of routes.the AdminAuth middleware function is used to enforce authentication for certain routes under the /admins group.The purpose of using the AdminAuth middleware here is to ensure that only authenticated administrators can access the routes under the /admins group.
	{
		// api.GET("/logout", authHandler.AdminLogout)
		// api.GET("/dashboard", authHandler.AdminDashboard)
		// api.GET("/sales-report", authHandler.FullSalesReport)

		//user management
		// userManagement := api.Group("/users")
		// {
		// 	userManagement.GET("/", authHandler.ListAllUsers)
		// 	userManagement.GET("/:id", authHandler.FindUserByID)
		// 	userManagement.PUT("/:id/block", authHandler.BlockUser)
		// 	userManagement.PUT("/unblock/:id", authHandler.UnblockUser)
		// }

		//admin management
		adminManagement := api.Group("/admins")
		{
			adminManagement.POST("/", authHandler.CreateAdmin)
		}

		// //category management routes
		// categoryRoutes := api.Group("/categories")
		// {
		// 	categoryRoutes.POST("/", productHandler.CreateCategory)
		// 	categoryRoutes.GET("/", productHandler.ListAllCategories)
		// 	categoryRoutes.GET("/:id", productHandler.FindCategoryByID)
		// 	categoryRoutes.PUT("/", productHandler.UpdateCategory)
		// 	categoryRoutes.DELETE("/:id", productHandler.DeleteCategory)
		// }

		// //brand management routes
		// brandRoutes := api.Group("/brands")
		// {
		// 	brandRoutes.POST("/", productHandler.CreateBrand)

		// }

		// //product management routes
		// productRoutes := api.Group("/products")
		// {
		// 	productRoutes.POST("/", productHandler.CreateProduct)
		// 	productRoutes.GET("/", productHandler.ListAllProducts)
		// 	productRoutes.GET("/:id", productHandler.FindProductByID)
		// 	productRoutes.PUT("/", productHandler.UpdateProduct)
		// 	productRoutes.DELETE("/:id", productHandler.DeleteProduct)
		// }

		// //product details routes
		// productDetails := api.Group("/product-details")
		// {
		// 	productDetails.POST("/", productHandler.AddProductDetails)

		// }

		// //coupon routes
		// couponRoutes := api.Group("/coupons")
		// {
		// 	couponRoutes.POST("/add", couponHandler.AddCoupon)
		// }

		// //order routes
		// orderRoutes := api.Group("/orders")
		// {
		// 	orderRoutes.PATCH("/update", orderHandler.UpdateOrderStatuses)
		// }

	}

}

// signUp := api.Group("/admin")
// 	{
// 		signUp.POST("/signup", adminHandler.CreateAdmin)
// 	}
// 	login := api.Group("/admin")
// 	{
// 		login.POST("/login", adminHandler.AdminLogin)
// 	}
