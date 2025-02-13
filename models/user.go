package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"size:100;not null"`
	Email     string         `gorm:"size:100;uniqueIndex;not null"`
	Password  string         `gorm:"size:255;not null"`
	CreatedAt int64          `gorm:"autoCreateTime"`
	UpdatedAt int64          `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
