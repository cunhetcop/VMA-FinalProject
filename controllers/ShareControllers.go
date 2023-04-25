package controllers

import (
	"fmt"
	"net/http"
	"nguyenhalinh/go/services"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type RegisterInput struct {
	ID              uint   `json:"-"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Images          string `json:"images" validate:"required"`
	Phone           string `json:"phone" validate:"required,numeric"`
	RoleID          uint   `json:"-"`
}
type LoginInput struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
type UpdateProfileInput struct {
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Images          string `json:"images" validate:"required"`
	Phone           string `json:"phone" validate:"required,numeric"` // Phone có thể để trống thì thay required bằng omitempty
	Password        string `json:"password,omitempty" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
}
type ProductInput struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Quantity    int     `json:"quantity" validate:"required,gt=0"`
	Images      string  `json:"images" validate:"required"`
	CategoryID  uint    `json:"category_id" validate:"required"`
	UserID      uint    `json:"user_id" binding:"required"`
}
type ForgotPasswordInput struct {
	Email string `json:"email" validate:"required,email"`
}


func getCurrentUserID(c *gin.Context) uint {

	tokenValue, exists := c.Get("token")

	if !exists {
		return 0
	}

	tokenString, ok := tokenValue.(string)

	if !ok {
		return 0
	}

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		id := uint(claims["user_id"].(float64))
		return id
	}

	return 0
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
		return
	}

	authParts := strings.Split(authHeader, " ")
	if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		return
	}

	token := authParts[1]

	expiration := 24 * time.Hour
	err := services.RevokeToken(token, expiration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}


func ForgotPassword(c *gin.Context) {
    var input services.ForgotPasswordInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    _, err := services.ResetPassword(&input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Chỉ trả về thông báo thành công mà không hiển thị mật khẩu mới
    c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("New password has been sent to %s", input.Email)})
}

