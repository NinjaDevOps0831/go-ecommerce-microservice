package handlerutil

import (
	"fmt"

	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int, error) {
	ID := c.Value("userID") //the value of userID is taken from the middleware (look the middleware, there the userID is retrieved from the jwttoken)
	fmt.Println("in handlerutil user id is", ID)
	userID, err := strconv.Atoi(fmt.Sprintf("%v", ID))
	return userID, err
}

// //both c.Get and c.Value can be used to retrieve data from context.. here i just used c.Value and i have also used the function GetUserIdFromContext
// id, valuebool := c.Get("userID")
// userid, err := strconv.Atoi(fmt.Sprintf("%v", id))
