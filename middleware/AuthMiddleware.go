// middleware/Auth_Middleware.go
package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"nguyenhalinh/go/database"
	"nguyenhalinh/go/models"
	"nguyenhalinh/go/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

const (
	JWT_SECRET = "linhnh4@vmodev.com" 
)


func GenerateToken(user models.User) (string, error) { 
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ 
		"user_id":  user.ID,
		"role_id":  user.RoleID,
		"exp": time.Now().Add(time.Hour * 72).Unix(), 
	})


	tokenString, err := token.SignedString([]byte(JWT_SECRET)) 
	if err != nil {
		return "", err 
	}

	return tokenString, nil
}


func ParseToken(tokenString string) (*jwt.Token, error) { 
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) { 
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}



func AuthMiddleware(redisClient *redis.Client) gin.HandlerFunc { 
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" { 
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || strings.ToLower(authParts[0]) != "bearer" { 
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}


		token, err := ParseToken(authParts[1]) 
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}


		revoked, err := utils.RedisClient.Get(utils.RedisClient.Context(), authParts[1]).Result()
		if err != redis.Nil && revoked == "true" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token revoked"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims) 
		userId := uint(claims["user_id"].(float64))
		var user models.User
		if err := database.DB.Preload("Role").Where("id = ?", userId).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", user) 
		c.Set("token", authParts[1])
		c.Next()
	}
}


