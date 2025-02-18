package models

import "github.com/gofrs/uuid"

type Apartment struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name        string    `gorm:"size:100;" json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Price       float64   `json:"price"`
	Available   bool      `json:"available"`
	Bookings    []Booking `gorm:"foreignKey:ApartmentID"`
}
