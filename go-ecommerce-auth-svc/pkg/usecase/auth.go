package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/config"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
	interfaces "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/repository/interface"
	services "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/usecase/interface"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	authRepo interfaces.AuthRepository
}

func NewUserUseCase(repo interfaces.AuthRepository) services.AuthUseCase {
	return &authUseCase{
		authRepo: repo,
	}

}

func (c *authUseCase) UserSignUp(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error) {
	checkUser, err := c.authRepo.FindUser(ctx, newUser)
	if err != nil {
		return response.UserDataOutput{}, err
	}

	//if that user not exists then create new user
	if checkUser.ID == 0 {
		//hash the pasword
		hashPasswd, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
		if err != nil {
			return response.UserDataOutput{}, errors.New("failed to hash the password")
		}
		newUser.Password = string(hashPasswd)

		userData, err := c.authRepo.UserSignUp(ctx, newUser)
		return userData, err
	}
	err = utils.CompareUsers(newUser, checkUser)
	return response.UserDataOutput{}, err

}

// Manage the otp verify status of users
func (c *authUseCase) OTPVerifyStatusManage(ctx context.Context, otpsession domain.OTPSession) error {
	err := c.authRepo.OTPVerifyStatusManage(ctx, otpsession)
	return err
}

// user login
func (c *authUseCase) LoginWithEmail(ctx context.Context, user request.UserLoginEmail) (domain.Users, error) {

	//dbUser, dberr := c.userRepo.FindUser(ctx, user)
	dbUser, dberr := c.authRepo.FindByEmail(ctx, user.Email)

	//check wether the user is found or not
	if dberr != nil {
		return dbUser, dberr
	} else if dbUser.ID == 0 {
		return dbUser, errors.New("user not exist with this details")
	}

	// check the user block_status to check wether user is blocked or not
	// if dbUser.BlockStatus {
	// 	return user, errors.New("user blocked by admin")
	// }

	userId := dbUser.ID

	blockStatus, err := c.authRepo.BlockStatus(ctx, userId)
	if blockStatus {
		return dbUser, errors.New("The user is blocked.. Please contact support")
	}

	if err != nil {
		return dbUser, err
	}

	//check the user password with dbPassword
	if bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)) != nil {
		return dbUser, errors.New("The entered password is wrong")
	}

	return dbUser, nil
}

//----user address

func (c *authUseCase) AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (domain.UserAddress, error) {

	address, err := c.authRepo.AddAddress(ctx, userAddressInput, userID)
	fmt.Println("debug check 1 usecase - address is", address, "err is", err)
	return address, err

}

func (c *authUseCase) CreateAdmin(ctx context.Context, newAdmin request.NewAdminInfo, adminID uint) (domain.Admin, error) {
	isSuperAdmin, err := c.authRepo.IsSuperAdmin(ctx, adminID)
	if err != nil {
		return domain.Admin{}, err
	}
	if !isSuperAdmin {
		return domain.Admin{}, fmt.Errorf("Only superadmin can create the new admins")
	}

	//hashing the admin password
	hash, err := bcrypt.GenerateFromPassword([]byte(newAdmin.Password), 10)
	if err != nil {
		return domain.Admin{}, err
	}
	newAdmin.Password = string(hash)
	fmt.Println("debug test -2 - newadmin.pasword is", newAdmin.Password)
	newAdminOutput, err := c.authRepo.CreateAdmin(ctx, newAdmin)
	return newAdminOutput, err

}
func (c *authUseCase) AdminLogin(ctx context.Context, input request.AdminLoginInfo) (string, response.AdminDataOutput, error) {
	var adminDataInModel response.AdminDataOutput
	// Now find the admindata with the given email from the databse
	adminInfo, err := c.authRepo.FindAdmin(ctx, input.Email)
	if err != nil {
		return "", adminDataInModel, fmt.Errorf("Error finding the admin")
	}
	if adminInfo.Email == "" {
		return "", adminDataInModel, fmt.Errorf("No such admin was found")

	}

	//Now compare and bcrypt the password
	if err := bcrypt.CompareHashAndPassword([]byte(adminInfo.Password), []byte(input.Password)); err != nil {
		return "", adminDataInModel, err
	}

	//Now check whether this admin is blocked by superadmin
	if adminInfo.IsBlocked {
		return "", adminDataInModel, fmt.Errorf("admin account is blocked")

	}

	// Now create JWT token and send it in cookie
	claims := jwt.MapClaims{
		"id":  adminInfo.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println("login debug 1- admin login - usecase - getconfig.jwt", []byte(config.GetConfig().JWT))

	tokenString, err := token.SignedString([]byte(config.GetConfig().JWT))

	fmt.Println("login debug 2- admin login - usecase - tokenstring", tokenString)
	//send back the created token

	//copying admin data from admindomain(from database table) to admin model struct
	// adminDataInModel.ID, adminDataInModel.UserName, adminDataInModel.Email, adminDataInModel.Phone, adminDataInModel.IsSuperAdmin = adminInfo.ID, adminInfo.UserName, adminInfo.Email, adminInfo.Phone, adminInfo.IsSuperAdmin   //t is a straightforward and concise method when you have a small number of fields to copy. However, it requires manually mapping each field, which can become cumbersome and error-prone if you have many fields or complex structures.
	copier.Copy(&adminDataInModel, &adminInfo) //Instead of using the above line for copying, we can use copier..  This method provides a more automated and flexible way of copying fields, especially when dealing with structs with a large number of fields or complex nested structures. The library handles field mapping based on struct tags, such as json, reducing the manual effort required.
	return tokenString, adminDataInModel, err
}
