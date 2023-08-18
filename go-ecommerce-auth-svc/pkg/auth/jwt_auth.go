package auth

import (
	"fmt"
	"time"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/config"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(id uint) (string, error) {
	//expireTime := time.Now().Add(60*time.Minute).Unix()

	//generate a jwt token
	// Create a new token object, specifying signing method and the claims you would like it to contain.
	//This creates a new JWT (JSON Web Token) with the specified claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	//This signs the token using the specified secret key and returns a string representation of the complete, signed token.
	tokenString, err := token.SignedString([]byte(config.GetConfig().JWT))

	fmt.Println("login debug 3- admin login - generate jwt - tokenstring", tokenString)
	if err != nil {
		return "", err
	}
	//return map[string]string{"accessToken": tokenString}, nil
	return tokenString, nil
}
