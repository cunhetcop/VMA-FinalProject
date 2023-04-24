// middleware/RBAC_Middleware.go
package middleware

import (
	"net/http"
	"strconv"

	"nguyenhalinh/go/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Permissions map[string][]string 
func RBACMiddleware(permissions Permissions, redisClient *redis.Client) gin.HandlerFunc {
    return func(c *gin.Context) {
        user, exists := c.Get("user")
        if !exists {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
            return
        }


        userRoleID := user.(models.User).RoleID



        allowedRoutes, exists := permissions[strconv.Itoa(int(userRoleID))]
        if !exists {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Role not found in permissions"})
            return
        }

        routeName := c.FullPath()
        isAllowed := false 
        for _, allowedRoute := range allowedRoutes {
            if allowedRoute == routeName {
                isAllowed = true
                break
            }
        }

        if !isAllowed { 
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denied"})
            return
        }

        c.Next()
    }
}

