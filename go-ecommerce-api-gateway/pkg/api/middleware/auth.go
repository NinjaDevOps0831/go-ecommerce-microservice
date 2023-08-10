package middleware

import (
	"fmt"
	"time"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-api-gateway/pkg/config"
	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// validate token for admin,, please merge this with the previous ValidateToken(for user) lateron
func ValidateToken2(tokenString string) (int, error) {
	//parses, validates, verifies the signature and returns the parsed token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		//retrieve the secret key which is stored in the env signing the string
		return []byte(config.GetJWTConfig()), nil
	})
	if err != nil {
		return 0, err
	}

	//extract the id claim from the token
	var parsedID interface{}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		parsedID = claims["id"]
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return 0, fmt.Errorf("token expired, re-login")
		}
	}
	//type assertion - It attempts to assert that the value stored in the parsedID variable is of type float64.
	value, ok := parsedID.(float64)
	if !ok {
		return 0, fmt.Errorf("expected a float value, but got %T, parsing id failed(error in middleware.auth)", parsedID)

	}

	id := int(value)
	return id, nil

}
