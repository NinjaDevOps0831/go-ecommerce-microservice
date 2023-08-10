package middleware

import (
	"fmt"
	"net/http"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/model/response"
	"github.com/gin-gonic/gin"
)

func AdminAuth(c *gin.Context) {
	tokenString, err := c.Cookie("AdminAuth")
	if err == http.ErrNoCookie {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "failed to login1", "AdminAuth cookie not present", nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(500, "failed to login - internal server error", err.Error(), nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	}
	fmt.Println("check2")
	adminID, err := ValidateToken2(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse(400, "failed to login2", err.Error(), nil))
		c.Abort() // Stop the execution of subsequent handlers
		return
	}
	c.Set("adminID", adminID)
	c.Next()

}
