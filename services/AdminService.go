// services/AdminService.go
package services

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"nguyenhalinh/go/database"
	"nguyenhalinh/go/middleware"
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/utils"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/crypto/bcrypt"
)

func CreateCategoryHandler(name string) (*models.Category, error) {
	category := models.Category{Name: name}
	if err := database.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdateCategoryHandler(categoryID uint, newName string) (*models.Category, error) {
	var category models.Category
	if err := database.DB.Where("id = ?", categoryID).First(&category).Error; err != nil {
		return nil, errors.New("category not found")
	}

	if err := database.DB.Model(&category).Update("name", newName).Error; err != nil {
		return nil, errors.New("error updating category")
	}

	return &category, nil
}

func DeleteCategoryHandler(categoryID uint) error {
	var category models.Category
	if err := database.DB.Where("id = ?", categoryID).First(&category).Error; err != nil {
		return errors.New("category not found")
	}

	if err := database.DB.Unscoped().Delete(&category).Error; err != nil {
		return errors.New("error deleting category")
	}

	return nil
}

func GetCategoryListAdminHandler() ([]models.Category, error) {
	var categories []models.Category
	if err := database.DB.Find(&categories).Error; err != nil {
		return nil, errors.New("error retrieving categories")
	}

	return categories, nil
}

func GetCategoryByIDAdminHandler(categoryID string) (models.Category, error) {
	var category models.Category
	if err := database.DB.First(&category, categoryID).Error; err != nil {
		return models.Category{}, errors.New("category not found")
	}

	return category, nil
}

//services
func CreateProductHandler(product *models.Product) error {
	if err := database.DB.Create(product).Error; err != nil {
		return errors.New("error creating product")
	}

	if err := database.DB.Preload("Category").First(product, product.ID).Error; err != nil {
		return errors.New("error fetching category")
	}

	return nil
}




func UpdateProductHandler(productID uint, updatedProduct *models.Product) error {
	var product models.Product
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return errors.New("product not found")
	}

	if err := database.DB.Model(&product).Updates(models.Product{
		Name:        updatedProduct.Name,
		Description: updatedProduct.Description,
		Price:       updatedProduct.Price,
		CategoryID:  updatedProduct.CategoryID,
		Images:      updatedProduct.Images,
		UserID:      updatedProduct.UserID,
	}).Error; err != nil {
		return errors.New("error updating product")
	}
	// Load Category for the updated product
	if err := database.DB.Preload("Category").First(&product, product.ID).Error; err != nil {
		return errors.New("error fetching category")
	}

	// Copy the updated product with the loaded Category to updatedProduct
	*updatedProduct = product
	return nil
}

func DeleteProductHandler(productID uint) error {
	var product models.Product
	if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
		return errors.New("product not found")
	}

	if err := database.DB.Unscoped().Delete(&product).Error; err != nil {
		return errors.New("error deleting product")
	}

	return nil
}

func GetProductListAdminHandler() ([]models.Product, error) {
	var products []models.Product
	if err := database.DB.Preload("Category").Preload("User").Find(&products).Error; err != nil {
		return nil, errors.New("error fetching products")
	}

	return products, nil
}

func GetProductByIDAdminHandler(productID uint) (models.Product, error) {
	var product models.Product
	if err := database.DB.Preload("Category").Preload("User").Where("id = ?", productID).First(&product).Error; err != nil {
		return models.Product{}, errors.New("product not found")
	}

	return product, nil
}

func RegisterAdminHandler(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("error hashing password")
	}

	adminRoleID, err := utils.GetRoleIDByName("admin")
	if err != nil {
		return errors.New("error retrieving admin role")
	}

	user.Password = string(hashedPassword)
	user.RoleID = adminRoleID

	if err := database.DB.Create(user).Error; err != nil {
		return errors.New("error registering admin")
	}

	return nil
}

func LoginAdminHandler(email, password string) (string, *models.User, error) {
	var admin models.User
	if err := database.DB.Preload("Role").Where("email = ?", email).First(&admin).Error; err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if admin.RoleID != models.AdminRoleID {
		return "", nil, errors.New("unauthorized")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(admin)
	if err != nil {
		return "", nil, errors.New("error generating token")
	}

	return token, &admin, nil
}

func GetUserListAdminHandler() ([]models.User, error) {
	var users []models.User
	if err := database.DB.Preload("Role").Find(&users).Error; err != nil {
		return nil, errors.New("error fetching users")
	}

	return users, nil
}

func GetUserAdminHandler(id string) (*models.User, error) {
	var user models.User
	if err := database.DB.Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func UpdateUserAdminHandler(currentUserID, targetUserID uint, firstName, lastName, images, phone, password string) (*models.User, error) {
	var currentUser, targetUser models.User
	if err := database.DB.Preload("Role").Where("id = ?", currentUserID).First(&currentUser).Error; err != nil {
		return nil, errors.New("current user not found")
	}

	if err := database.DB.Preload("Role").Where("id = ?", targetUserID).First(&targetUser).Error; err != nil {
		return nil, errors.New("target user not found")
	}

	if currentUser.ID != targetUser.ID && targetUser.RoleID == 2 {
		return nil, errors.New("cannot update other admin accounts")
	}

	targetUser.FirstName = firstName
	targetUser.LastName = lastName
	targetUser.Images = images
	targetUser.Phone = phone

	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("error hashing the new password")
		}
		targetUser.Password = string(hashedPassword)
	}

	if err := database.DB.Save(&targetUser).Error; err != nil {
		return nil, errors.New("error updating user")
	}

	return &targetUser, nil
}


func DeleteUserAdminHandler(userID, currentUserID uint, token string) error {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.RoleID == 2 && currentUserID != user.ID {
		return errors.New("cannot delete other admin accounts")
	}

	if err := database.DB.Unscoped().Delete(&user).Error; err != nil {
		return errors.New("error deleting user")
	}

	if currentUserID == user.ID {
		expiration := 24 * time.Hour
		err := RevokeToken(token, expiration)
		if err != nil {
			return errors.New("failed to revoke token")
		}
	}

	return nil
}

func UploadUserImageHandler(userID, currentUserID uint, fileReader io.Reader, imagePath string) error {
	var user models.User
	if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return errors.New("user not found")
	}

	if user.RoleID == 2 && currentUserID != user.ID {
		return errors.New("cannot upload image for other admin accounts")
	}

	// Khởi tạo session và S3 client
	sess := utils.InitAWSSession()
	s3Client := utils.InitS3Client(sess)

// Đọc tệp và chuyển đổi nó thành đối tượng io.Reader
buf := new(bytes.Buffer)
_, err := buf.ReadFrom(fileReader)
if err != nil {
	return errors.New("error reading file")
}

// Tính toán MD5 hash của nội dung tệp
hasher := md5.New()
_, err = io.Copy(hasher, bytes.NewReader(buf.Bytes())) // Sửa lại dòng này
if err != nil {
	return errors.New("error calculating MD5 hash")
}
contentMD5 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

// Xác định loại nội dung của tệp
contentType := http.DetectContentType(buf.Bytes())

// Tải tệp lên S3
input := &s3.PutObjectInput{
	Body:        aws.ReadSeekCloser(bytes.NewReader(buf.Bytes())), // Sửa lại dòng này
	Bucket:      aws.String("vma-demo"), // Đã thay đổi tên thùng chứa S3
	Key:         aws.String(imagePath),
	ContentMD5:  aws.String(contentMD5),
	ContentType: aws.String(contentType), // Thêm dòng này để đặt loại nội dung
}

_, err = s3Client.PutObject(input)
if err != nil {
	fmt.Printf("S3 Error Detail: %v\n", err) // Thêm dòng này để in thông tin chi tiết của lỗi
	return errors.New("error uploading image to S3")
}
baseURL := "https://vma-demo.s3.ap-southeast-1.amazonaws.com/"
imageURL := baseURL + imagePath
user.Images = imageURL

if err := database.DB.Save(&user).Error; err != nil {
	return errors.New("error updating user with new image")
}

return nil
}

func UploadProductImageHandler(productID uint, fileReader io.Reader, imagePath string) error {
	var product models.Product
    if err := database.DB.Where("id = ?", productID).First(&product).Error; err != nil {
        return errors.New("product not found")
    }

    sess := utils.InitAWSSession()
    s3Client := utils.InitS3Client(sess)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(fileReader)
	if err != nil {
		return errors.New("error reading file")
	}

	hasher := md5.New()
	_, err = io.Copy(hasher, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return errors.New("error calculating MD5 hash")
	}
	contentMD5 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	contentType := http.DetectContentType(buf.Bytes())

	input := &s3.PutObjectInput{
		Body:        aws.ReadSeekCloser(bytes.NewReader(buf.Bytes())),
		Bucket:      aws.String("vma-demo"),
		Key:         aws.String(imagePath),
		ContentMD5:  aws.String(contentMD5),
		ContentType: aws.String(contentType),
	}

	_, err = s3Client.PutObject(input)
	if err != nil {
		fmt.Printf("S3 Error Detail: %v\n", err)
		return errors.New("error uploading image to S3")
	}

	baseURL := "https://vma-demo.s3.ap-southeast-1.amazonaws.com/"
	imageURL := baseURL + imagePath
	product.Images = imageURL

	if err := database.DB.Save(&product).Error; err != nil {
		return errors.New("error updating product with new image")
	}

	return nil
}