// routers/user_router.go
package routers

import (
	"nguyenhalinh/go/controllers"
	"nguyenhalinh/go/middleware"

	"nguyenhalinh/go/utils"

	"github.com/gin-gonic/gin"
)


func SetupUserRoutes(r *gin.Engine, userPermissions middleware.Permissions) {
    userRoutes := r.Group("/user")
    {
        userRoutes.POST("/register", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
        userRoutes.GET("/forgot-password", controllers.ForgotPassword)
        userRoutes.Use(middleware.AuthMiddleware(utils.RedisClient))
        userRoutes.Use(middleware.RBACMiddleware(userPermissions, utils.RedisClient))
        userRoutes.GET("/logout", controllers.Logout)
        userRoutes.GET("/profile", controllers.GetMyProfile)
        userRoutes.PUT("/profile", controllers.UpdateMyProfile)
        userRoutes.DELETE("/profile", controllers.DeleteMyProfile)
        
        // New routes for products and categories
        userRoutes.GET("/products", controllers.GetProductList)
        userRoutes.GET("/products/:id", controllers.GetProductByID)
        userRoutes.GET("/categories", controllers.GetCategoryList)
        userRoutes.GET("/categories/:id", controllers.GetCategorybyID)
    }
}
