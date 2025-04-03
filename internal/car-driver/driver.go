package carDriver

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID             uuid.UUID `gorm:"column:id" json:"id"`
	Name           string    `gorm:"column:name" json:"name"`
	PhoneNumber    string    `gorm:"column:phone_number" json:"phone_number"`
	Email          string    `gorm:"column:email" json:"email"`
	Address        string    `gorm:"column:address" json:"address"`
	Identification string    `gorm:"column:identification" json:"identification"`
	LicenseNumber  string    `gorm:"column:license_number" json:"license_number"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}
