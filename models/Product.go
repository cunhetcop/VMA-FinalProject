// models/Product.go
package models

import (
	"time"

	"gorm.io/gorm"
)
type Product struct {
    ID        uint   `gorm:"primaryKey" json:"id"`
    Name        string  `gorm:"not null" validate:"required" json:"name"`
    Description string  `gorm:"not null" validate:"required" json:"description"`
    Price       float64 `gorm:"not null" validate:"required,gt=0" json:"price"`
    Quantity    int     `gorm:"not null" validate:"required,gt=0" json:"quantity"`
    Images      string  `gorm:"not null" validate:"required" json:"images"`
    UserID      uint    `gorm:"not null" json:"user_id"`
    User        User    `gorm:"foreignKey:UserID" json:"-"`
    CategoryID  uint    `gorm:"not null" json:"-"`
    Category    Category `gorm:"foreignKey:CategoryID" validate:"required" json:"category"`
    CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}



