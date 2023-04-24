// models/Category.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"unique;not null" validate:"required" json:"Name"`
	CreatedAt time.Time `gorm:"not null" json:"-"`
	UpdatedAt time.Time `gorm:"not null" json:"-"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
