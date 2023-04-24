// models/User.go
package models

import (
	"time"

	"gorm.io/gorm"
)

const AdminRoleID = 2 
type User struct {
    ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
    FirstName string `gorm:"not null" validate:"required" json:"first_name"`
    LastName  string `gorm:"not null" validate:"required" json:"last_name"`
    Email     string `gorm:"unique;not null" validate:"required,email" json:"email"`
    Password  string `gorm:"not null" validate:"required,min=8" json:"-"`
    Images    string `gorm:"not null" validate:"required" json:"images"`
    Phone     string `gorm:"not null" validate:"required,numeric" json:"phonenumber"`
    RoleID    uint   `gorm:"not null" json:"-"`
    Role      Role   `gorm:"not null" json:"role"`
    CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

