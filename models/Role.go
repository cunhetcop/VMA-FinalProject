// models/Role.go
package models

import (
	"time"

	"gorm.io/gorm"
)
type Role struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name        string `gorm:"unique;not null" validate:"required" json:"name"`
    CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

