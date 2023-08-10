package handlerutil

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAdminIdFromContext(c *gin.Context) (int, error) {
	id := c.Value("adminID") //the value of adminID is taken from the middleware (look the middleware, there the adminID is retrieved from the jwttoken)
	//fmt.Println("in handlerutil admin id is", id)
	adminID, err := strconv.Atoi(fmt.Sprintf("%v", id)) //fmt.Sprintf function to format the value of id as a string. strconv.Atoi: This function is used to convert the string value obtained from fmt.Sprintf to an integer.
	return adminID, err
}
