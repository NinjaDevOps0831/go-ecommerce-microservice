package repository

import (
	"context"
	"fmt"

	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/domain"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/request"
	"github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/model/response"
	interfaces "github.com/ajujacob88/go-ecommerce-microservice-clean-arch/go-ecommerce-auth-svc/pkg/repository/interface"
	"gorm.io/gorm"
)

type authDatabase struct {
	DB *gorm.DB
}

func NewAuthRepository(DB *gorm.DB) interfaces.AuthRepository {
	return &authDatabase{DB}
}

func (c *authDatabase) UserSignup(ctx context.Context, newUser request.NewUserInfo) (response.UserDataOutput, error) {
	var userData response.UserDataOutput

	//save the user details
	UserSignUpQuery := `INSERT INTO users(first_name, last_name, email, phone, password, created_at)
						VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id,first_name, last_name, email, phone`

	err := c.DB.Raw(UserSignUpQuery, newUser.FirstName, newUser.LastName, newUser.Email, newUser.Phone, newUser.Password).Scan(&userData).Error

	if err != nil {
		return response.UserDataOutput{}, fmt.Errorf("failed to create the user %s", newUser.FirstName)
	}

	//insert the data into userinfo table
	insertUserinfoQuery := `INSERT INTO user_infos (is_verified, is_blocked,users_id)
							VALUES ('f','f',$1);`
	err = c.DB.Exec(insertUserinfoQuery, userData.ID).Error
	if err != nil {
		return response.UserDataOutput{}, fmt.Errorf("failed to create the user(falied to copy to userinfo table) %s", newUser.FirstName)
	}

	return userData, err
}

func (c *authDatabase) FindByEmail(ctx context.Context, email string) (domain.Users, error) {
	//var user domain.Users
	// err := c.DB.Where("Email = ?", email).Find(&user).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return domain.Users{}, errors.New("invalid email")
	// 	}
	// 	return domain.Users{}, err
	// }
	var userData domain.Users
	fmt.Println("email is", email, " and users.email is")
	findUserQuery := `	SELECT users.id, users.first_name, users.last_name, users.email, users.phone, users.password 
						FROM users 
						WHERE users.email = $1;`

	err := c.DB.Raw(findUserQuery, email).Scan(&userData).Error
	fmt.Println("error is", err)
	// if userData.BlockStatus {
	// 	return userData, errors.New("you are blocked")
	// }
	fmt.Println("userdata is", userData)

	return userData, err
}

func (c *authDatabase) BlockStatus(ctx context.Context, userId uint) (bool, error) {

	blockStatusQuery := `SELECT is_blocked FROM user_infos WHERE users_id = $1;`

	var blockStatus bool

	err := c.DB.Raw(blockStatusQuery, userId).Scan(&blockStatus).Error
	return blockStatus, err
}

func (c *authDatabase) AddAddress(ctx context.Context, userAddressInput request.UserAddressInput, userID uint) (domain.UserAddress, error) {
	var addedAddress domain.UserAddress

	insertAddressQuery := `	INSERT INTO user_addresses(
								user_id, house_number, street, city, district, state, pincode, landmark) 
								VALUES($1,$2,$3,$4,$5,$6, $7, $8) RETURNING *`
	err := c.DB.Raw(insertAddressQuery, userID, userAddressInput.HouseNumber, userAddressInput.Street, userAddressInput.City, userAddressInput.District, userAddressInput.State, userAddressInput.Pincode, userAddressInput.Landmark).Scan(&addedAddress).Error

	if err != nil {
		return domain.UserAddress{}, err
	}
	return addedAddress, nil

}

func (c *authDatabase) CreateAdmin(ctx context.Context, newAdminInfo request.NewAdminInfo) (domain.Admin, error) {
	var newAdmin domain.Admin
	createAdminQuery := `	INSERT INTO admins(user_name, email, phone, password,is_super_admin, is_blocked, created_at, updated_at)
						 	VALUES($1, $2, $3,$4, false, false, NOW(), NOW()) RETURNING *;`
	err := c.DB.Raw(createAdminQuery, newAdminInfo.UserName, newAdminInfo.Email, newAdminInfo.Phone, newAdminInfo.Password).Scan(&newAdmin).Error
	newAdmin.Password = "" //By setting it to an empty string before returning, the function ensures that the password is not accessible outside of the function scope.
	return newAdmin, err
}

func (c *authDatabase) FindAdmin(ctx context.Context, email string) (domain.Admin, error) {
	var adminData domain.Admin
	findAdminQuery := `	SELECT *
						FROM admins
						WHERE email = $1;`

	err := c.DB.Raw(findAdminQuery, email).Scan(&adminData).Error
	return adminData, err
}
