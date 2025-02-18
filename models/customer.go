package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	FirstName   string    `gorm:"size:100;" json:"first_name"`
	LastName    string    `gorm:"size:100;" json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Email       string    `gorm:"size:100;" json:"email"`
	Phone       string    `gorm:"size:15;uniqueIndex;not null" json:"phone"`
	Bookings    []Booking `gorm:"foreignKey:CustomerID"`
}
