package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"nguyenhalinh/go/database"
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/services"
)

func CreateCategoryAdmin(c *gin.Context) {
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := services.CreateCategoryHandler(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
}

func UpdateCategoryAdmin(c *gin.Context) {
	var input struct {
		Name string `json:"name" validate:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	updatedCategory, err := services.UpdateCategoryHandler(uint(categoryID), input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully", "category": updatedCategory})
}

func DeleteCategoryAdmin(c *gin.Context) {
	categoryIDStr := c.Param("id")
	categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	err = services.DeleteCategoryHandler(uint(categoryID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func GetCategoryListAdmin(c *gin.Context) {
	categories, err := services.GetCategoryListAdminHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

func GetCategorybyIDAdmin(c *gin.Context) {
	categoryID := c.Param("id")

	category, err := services.GetCategoryByIDAdminHandler(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func CreateProductAdmin(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		Images:      input.Images,
		UserID:      input.UserID,
	}

	err := services.CreateProductHandler(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully", "product": product})
}

func UpdateProductAdmin(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	productID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	updatedProduct := models.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		Images:      input.Images,
		UserID:      input.UserID,
	}

	err = services.UpdateProductHandler(uint(productID), &updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully" , "product": updatedProduct})
}

func DeleteProductAdmin(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	err = services.DeleteProductHandler(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func GetProductListAdmin(c *gin.Context) {
	products, err := services.GetProductListAdminHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func GetProductAdmin(c *gin.Context) {
	productID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := services.GetProductByIDAdminHandler(uint(productID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func RegisterAdmin(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  input.Password,
		Images:    input.Images,
		Phone:     input.Phone,
	}

	err := services.RegisterAdminHandler(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Admin registered successfully"})
}

func LoginAdmin(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, admin, err := services.LoginAdminHandler(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "admin": admin})
}

// controllers/adminController.go
func GetUserListAdmin(c *gin.Context) {
	users, err := services.GetUserListAdminHandler()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func GetUserAdmin(c *gin.Context) {
	user, err := services.GetUserAdminHandler(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUserAdmin(c *gin.Context) {
	var input UpdateProfileInput
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDUint, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	currentUserID := getCurrentUserID(c)

	updatedUser, err := services.UpdateUserAdminHandler(currentUserID, uint(userIDUint), input.FirstName, input.LastName, input.Images, input.Phone, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": updatedUser})
}


func DeleteUserAdmin(c *gin.Context) {
	userIDUint, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	currentUserID := getCurrentUserID(c)
	token := c.GetString("token")

	err = services.DeleteUserAdminHandler(uint(userIDUint), currentUserID, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func UploadUsersImage(c *gin.Context) {
	userIDUint, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	currentUserID := getCurrentUserID(c)

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
		return
	}
	defer fileReader.Close()

	imageFileName := fmt.Sprintf("image_%d.png", time.Now().UnixNano())
	imagePath := filepath.Join("uploads", imageFileName)
    baseURL := "https://vma-demo.s3.ap-southeast-1.amazonaws.com/"
    imageURL := baseURL + imagePath
	err = services.UploadUserImageHandler(uint(userIDUint), currentUserID, fileReader, imagePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded and assigned to user successfully", "image_url": imageURL})
}

func UploadProductsImage(c *gin.Context) {
    productIDUint, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
        return
    }

    currentUserID := getCurrentUserID(c)

    // Kiểm tra quyền admin
    var currentUser models.User
    if err := database.DB.Where("id = ?", currentUserID).First(&currentUser).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }
    if currentUser.RoleID != models.AdminRoleID {
        c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to upload images for products"})
        return
    }

    file, err := c.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file"})
        return
    }

    fileReader, err := file.Open()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading file"})
        return
    }
    defer fileReader.Close()

    imageFileName := fmt.Sprintf("image_%d.png", time.Now().UnixNano())
    imagePath := filepath.Join("uploads", imageFileName)
    baseURL := "https://vma-demo.s3.ap-southeast-1.amazonaws.com/"
    imageURL := baseURL + imagePath

    err = services.UploadProductImageHandler(uint(productIDUint), fileReader, imagePath)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Image uploaded and assigned to product successfully", "image_url": imageURL})
}