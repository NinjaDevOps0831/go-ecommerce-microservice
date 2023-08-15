package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"golang.org/x/crypto/bcrypt"
)

func CompareUsers(newUser request.NewUserInfo, checkUser domain.Users) (err error) {
	if checkUser.Email == newUser.Email {
		err = errors.Join(err, errors.New("user already eists with this email"))

	}
	if checkUser.Phone == newUser.Phone {
		err = errors.Join(err, errors.New("user already exists with this phone number"))
	}
	return err
}

// generate coupon codes with the first 3 letters the same for all codes, the next 3 characters randomly generated, and the last 3 characters as generated numbers
func GenerateCouponCode() string {
	// Constants
	const prefix = "SMT" //  desired prefix
	const randomChars = 3
	const numberChars = 3

	// Pool of characters for random generation
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers := "0123456789"

	// Generate random part of the coupon code
	rand.Seed(time.Now().UnixNano())
	randomCode := make([]byte, randomChars)
	for i := range randomCode {
		randomCode[i] = letters[rand.Intn(len(letters))]
	}

	// Generate random part of the coupon code as numbers
	numberCode := make([]byte, numberChars)
	for i := range numberCode {
		numberCode[i] = numbers[rand.Intn(len(numbers))]
	}

	// Concatenate prefix, random code, and number code
	couponCode := prefix + string(randomCode) + string(numberCode)

	return couponCode
}

func StringToTime(timeString string) (time.Time, error) {

	// parse the string time to time
	const shortForm = "2006-Jan-02"
	timeValue, err := time.Parse(shortForm, timeString)

	if err != nil {
		return timeValue, fmt.Errorf("faild to parse given time %v to time variable \n invalid input", timeString)
	}
	return timeValue, err
}

// password hashing
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14) //GenerateFromPassword returns the bcrypt hash of the password
	return string(bytes), err
}
