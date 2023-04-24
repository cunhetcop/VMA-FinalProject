// routers/user_router.go
package routers

import (
	"net/http"
	"nguyenhalinh/go/controllers"
	"nguyenhalinh/go/middleware"

	"nguyenhalinh/go/utils"

	"github.com/gin-gonic/gin"
)

//hàm này để đăng ký các route cho user
func SetupUserRoutes(r *gin.Engine, userPermissions middleware.Permissions) { //truyền vào một router và một map có key là id của role và value là một slice các string là các đường dẫn mà role đó được phép truy cập
    userRoutes := r.Group("/user") //tạo một group route có prefix là /user
    {
        userRoutes.StaticFS("/uploads", http.Dir("uploads")) // etc http://localhost:8080/admin/uploads/image_1681259025825561000.png
        userRoutes.POST("/register", controllers.RegisterUser)
        userRoutes.POST("/login", controllers.LoginUser)
        userRoutes.GET("/logout", controllers.Logout)
        userRoutes.GET("/forgotpassword", controllers.ForgotPassword)
        userRoutes.Use(middleware.AuthMiddleware(utils.RedisClient))
        userRoutes.Use(middleware.RBACMiddleware(userPermissions, utils.RedisClient))
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
