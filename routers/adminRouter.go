// router/admin_router.go
package routers

import (
	"nguyenhalinh/go/controllers"
	"nguyenhalinh/go/middleware"
	"nguyenhalinh/go/utils"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine, adminPermissions middleware.Permissions) { 
    adminRoutes := r.Group("/admin")
    {
        adminRoutes.POST("/register", controllers.RegisterAdmin)
        adminRoutes.POST("/login", controllers.LoginAdmin)
        adminRoutes.GET("/forgot-password", controllers.ForgotPassword)
        adminRoutes.Use(middleware.AuthMiddleware(utils.RedisClient))
        adminRoutes.Use(middleware.RBACMiddleware(adminPermissions, utils.RedisClient))
        adminRoutes.GET("/logout", controllers.Logout)
        adminRoutes.GET("/users", controllers.GetUserListAdmin)
        adminRoutes.GET("/users/:id", controllers.GetUserAdmin)
        adminRoutes.PUT("/users/:id", controllers.UpdateUserAdmin)
        adminRoutes.DELETE("/users/:id", controllers.DeleteUserAdmin)
		adminRoutes.POST("/upload-avatar-user/:id", controllers.UploadUsersImage)

        // New routes for products and categories
        adminRoutes.GET("/products", controllers.GetProductListAdmin)
        adminRoutes.GET("/products/:id", controllers.GetProductAdmin)
        adminRoutes.POST("/products", controllers.CreateProductAdmin)
        adminRoutes.PUT("/products/:id", controllers.UpdateProductAdmin)
        adminRoutes.DELETE("/products/:id", controllers.DeleteProductAdmin)
        adminRoutes.POST("/upload-avatar-product/:id", controllers.UploadProductsImage)

        adminRoutes.GET("/categories", controllers.GetCategoryListAdmin)
        adminRoutes.GET("/categories/:id", controllers.GetCategorybyIDAdmin)
        adminRoutes.POST("/categories", controllers.CreateCategoryAdmin)
        adminRoutes.PUT("/categories/:id", controllers.UpdateCategoryAdmin)
        adminRoutes.DELETE("/categories/:id", controllers.DeleteCategoryAdmin)

	}
}
