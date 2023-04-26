// Package: main
package main

import (
	"nguyenhalinh/go/database"
	"nguyenhalinh/go/database/seeder"
	"nguyenhalinh/go/middleware"
	"nguyenhalinh/go/routers"
	"nguyenhalinh/go/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        panic("Error loading app.env file")
    }
    database.Connect()
    database.MigrationSchemasUp()
    utils.InitRedis()
    
    utils.ScheduleDeleteExpiredTokens()
    
    userRole, adminRole, err := seeder.CreateDefaultRoles()
    if err != nil {
        panic(err)
    }
    // utils.MigrationSchemasDown()
    r := gin.Default()

    userPermissions := middleware.Permissions{
        strconv.Itoa(int(userRole.ID)): []string{
            "/user/register",
            "/user/profile",
            "/user/products",
            "/user/products/:id",
            "/user/categories",
            "/user/categories/:id",
            "/user/logout",
            "/user/forgot-password",
        },
    }

    adminPermissions := middleware.Permissions{
        strconv.Itoa(int(adminRole.ID)): []string{
            "/admin/users",
            "/admin/users/:id",
            "/admin/products",
            "/admin/products/:id",
            "/admin/categories",
            "/admin/categories/:id",
            "/admin/upload-avatar-user/:id",
            "/admin/upload-avatar-product/:id",
            "/admin/logout",
            "/admin/forgot-password",
        },
    }

    routers.SetupUserRoutes(r, userPermissions)
    routers.SetupAdminRoutes(r, adminPermissions)
    r.Run(":8080")
}

