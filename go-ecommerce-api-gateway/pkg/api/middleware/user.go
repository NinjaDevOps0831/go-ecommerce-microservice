package middleware

import (
	"fmt"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/response"
	"github.com/gin-gonic/gin"
)

func UserAuth(c *gin.Context) {
	tokenString, err := c.Cookie("UserAuth")
	if err == http.ErrNoCookie { // this error occurs if cookie while login is not correctly set for eg here c.SetCookie("UserAuth", tokenString["accessToken"], 60*60, "", "", false, true) in login handler
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "failed to login1", "UserAuth cookie not present", nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(500, "failed to login - internal server error", err.Error(), nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	}

	fmt.Println("user token string is", tokenString)

	userID, err := ValidateToken2(tokenString)
	fmt.Println("check3 in login middleware userid is", userID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "failed to login user", err.Error(), nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	}
	c.Set("userID", userID)
	c.Next()
}
