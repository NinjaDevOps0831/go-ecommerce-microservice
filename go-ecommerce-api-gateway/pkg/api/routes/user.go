package routes

import (
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/handler"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(
	api *gin.RouterGroup,
	authHandler handler.AuthHandler,
	// productHandler handler.ProductHandler,
	// cartHandler handler.CartHandler,
	// orderHandler handler.OrderHandler,

) {

	// User routes that don't require authentication
	//sets up a route group for the "/user" endpoint
	signup := api.Group("/user")
	{
		signup.POST("/signup", authHandler.UserSignUp)
		signup.POST("/signup/otp/verify", authHandler.SignupOtpVerify)
	}

	// products := api.Group("/user")
	// {
	// 	products.GET("/products", productHandler.ListAllProducts)
	// }

	login := api.Group("/user")
	{
		login.POST("/login/email", authHandler.UserLoginByEmail)
	}

	// forgotPassword := api.Group("/user/forgot-password")
	// {
	// 	forgotPassword.POST("/", authHandler.ForgotPassword)
	// 	forgotPassword.PATCH("/otp-verify", authHandler.ForgotPasswordOtpVerify)
	// }

	//User routes that require authentication

	home := api.Group("/user")
	{
		//AuthorizationMiddleware as middleware to perform authorization checks for users accessing the "/user" endpoint.
		home.Use(middleware.UserAuth)
		{
			// home.GET("/home", authHandler.Homehandler)
			// home.GET("/logout", authHandler.LogoutHandler)
			// //home.GET("/products", productHandler.ListAllProducts)
			// home.GET("/products/:id", productHandler.FindProductByID)
			// home.POST("/cart/add/:product_details_id", cartHandler.AddToCart)
			// home.DELETE("/cart/remove/:product_details_id", cartHandler.RemoveFromCart)
			// home.GET("/cart", cartHandler.ViewCart)
			home.POST("/addresses", authHandler.AddAddress)
			// home.PATCH("/addresses/edit/:address_id", authHandler.UpdateAddress)
			// home.DELETE("/addresses/:address_id", authHandler.DeleteAddress)
			// home.GET("/addresses", authHandler.ListAddress)

			// //home.PATCH("/cart/applycoupon", couponHandler.ApplyCouponToCart)

			// home.POST("/cart/placeorder", orderHandler.PlaceOrderFromCart)

			// //	home.GET("/payments/razorpay/:order_id", paymentHandler.RazorpayCheckout)
			// //home.POST("/payments/success", paymentHandler.RazorpayVerify)

			// home.POST("/orders/return", orderHandler.ReturnRequest)

			// home.PATCH("/orders/cancel/:order_id", orderHandler.CancellOrder)

			// home.GET("/orders", orderHandler.ViewAllOrders)

			// home.GET("/coupons", couponHandler.ViewAllCoupons)

		}
	}

}

// api.POST("/signup", userHandler.UserSignUp)
// api.POST("/signup/otp/verify", userHandler.SignupOtpVerify)
// api.POST("/login/email", userHandler.UserLoginByEmail)

// api.Use(middleware.AuthorizationMiddleware("user"))
// {
// 	api.GET("/home", userHandler.UserProfile)
// 	api.GET("/logout", userHandler.UserLogout)

// }
