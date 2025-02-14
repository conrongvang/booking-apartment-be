package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string         `gorm:"size:100;" json:"first_name"`
	LastName  string         `gorm:"size:100;" json:"last_name"`
	Email     string         `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"size:255;not null" json:"password"`
	Phone     string         `gorm:"size:15;" json:"phone"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
