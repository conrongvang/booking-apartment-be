package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Booking struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	CustomerID  uuid.UUID      `gorm:"type:uuid;not null" json:"customer_id"`
	ApartmentID uuid.UUID      `gorm:"type:uuid;not null" json:"apartment_id"`
	CheckIn     time.Time      `gorm:"type:timestamp" json:"checkin"`
	CheckOut    time.Time      `gorm:"type:timestamp" json:"checkout"`
	TotalAmount float64        `gorm:"type:float" json:"total_amount"`
	CreatedAt   int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`

	Customer  Customer  `gorm:"foreignKey:CustomerID"`
	Apartment Apartment `gorm:"foreignKey:ApartmentID"`
}
