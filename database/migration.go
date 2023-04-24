// database/migration.go
package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrationSchemasUp() {
    // Lấy đường dẫn tuyệt đối của thư mục hiện tại (chứa file migration.go)
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to get current directory: %v", err)
    }
    migrationsPath := filepath.Join(currentDir, "database", "migrations")

 // Trích xuất *sql.DB từ gorm.DB
 sqlDB, err := DB.DB()
 if err != nil {
	 log.Fatalf("Failed to extract *sql.DB from gorm.DB: %v", err)
 }

 // Tạo một instance mới của postgres.Driver
 postgresDriver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
 if err != nil {
	 log.Fatalf("Failed to create postgres driver: %v", err)
 }

 // Tạo một instance mới của migrate.Migrate và liên kết nó với các file migration và driver postgres
 m, err := migrate.NewWithDatabaseInstance(
	 "file://"+migrationsPath, 
	 "postgres",               
	 postgresDriver,           
 )
 if err != nil {
	 log.Fatalf("Failed to create migration instance: %v", err)
 }

 //up migrations
 err = m.Up()
 if err != nil && err != migrate.ErrNoChange {
	 log.Fatalf("Failed to run migrations: %v", err)
 } else {
	 log.Println("Migrations run successfully")
 }
}


func MigrationSchemasDown() {
    // Lấy đường dẫn tuyệt đối của thư mục hiện tại (chứa file migration.go)
    currentDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Failed to get current directory: %v", err)
    }
    migrationsPath := filepath.Join(currentDir, "database", "migrations")

 // Trích xuất *sql.DB từ gorm.DB
 sqlDB, err := DB.DB()
 if err != nil {
	 log.Fatalf("Failed to extract *sql.DB from gorm.DB: %v", err)
 }

 // Tạo một instance mới của postgres.Driver
 postgresDriver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
 if err != nil {
	 log.Fatalf("Failed to create postgres driver: %v", err)
 }

 // Tạo một instance mới của migrate.Migrate và liên kết nó với các file migration và driver postgres
 m, err := migrate.NewWithDatabaseInstance(
	 "file://"+migrationsPath, 
	 "postgres",               
	 postgresDriver,           
 )
 if err != nil {
	 log.Fatalf("Failed to create migration instance: %v", err)
 }

    //down migrations
    err = m.Down()
    if err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Failed to rollback migrations: %v", err)
    } else {
        log.Println("Migrations rolled back successfully")
    }
}
