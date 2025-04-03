package carDriver

import (
	"time"

	"github.com/google/uuid"
)

type Driver struct {
	ID             uuid.UUID `gorm:"column:id"`
	Name           string    `gorm:"column:name"`
	PhoneNumber    string    `gorm:"column:phone_number"`
	Email          string    `gorm:"column:email"`
	Address        string    `gorm:"column:address"`
	Identification string    `gorm:"column:identification"`
	LicenseNumber  string    `gorm:"column:license_number"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
