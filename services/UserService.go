// services/userService.go
package services

import (
	"errors"
	"nguyenhalinh/go/database"
	"nguyenhalinh/go/middleware"
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserHandler(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	userRoleID, err := utils.GetRoleIDByName("user")
	if err != nil {
		return errors.New("error retrieving user role")
	}

	user.Password = string(hashedPassword)
	user.RoleID = userRoleID

	if err := database.DB.Create(user).Error; err != nil {
		return errors.New("error registering user")
	}

	return nil
}

// services/userService.go
func LoginUserHandler(email, password string) (string, *models.User, error) {
	var dbUser models.User
	if err := database.DB.Preload("Role").Where("email = ?", email).First(&dbUser).Error; err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(dbUser)
	if err != nil {
		return "", nil, errors.New("error generating token")
	}

	return token, &dbUser, nil
}

func GetMyProfileHandler(user *models.User) *models.User {
	// Không cần thực hiện các thao tác cần thiết để lấy thông tin người dùng từ cơ sở dữ liệu
	// Vì bạn đã có thông tin người dùng từ gin.Context

	// Trả về thông tin người dùng
	return user
}

func UpdateProfileHandler(user *models.User, firstName, lastName, images, phone, password string) error {
	user.FirstName = firstName
	user.LastName = lastName
	user.Images = images
	user.Phone = phone

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return errors.New("error hashing new password")
		}
		user.Password = string(hashedPassword)
	}

	if err := database.DB.Preload("Role").Save(user).Error; err != nil {
		return errors.New("error updating user profile")
	}

	return nil
}

func DeleteUser(user *models.User, token string) error {
	err := database.DB.Unscoped().Delete(user).Error
	if err != nil {
		return errors.New("error deleting user profile")
	}

	expiration := 24 * time.Hour
	err = RevokeToken(token, expiration)
	if err != nil {
		return errors.New("failed to revoke token")
	}

	return nil
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Preload("Category").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductByID(id string) (*models.Product, error) {
	var product models.Product
	if err := database.DB.Preload("Category").First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func GetCategoryList() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func GetCategoryByID(id string) (*models.Category, error) {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func UploadImage(userID uint, encodedImage string) error {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&user).Update("Images", encodedImage).Error; err != nil {
		return err
	}

	return nil
}
