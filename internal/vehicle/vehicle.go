package vehicle

import (
	"time"

	"github.com/google/uuid"
)

type Vehicle struct {
	ID          uuid.UUID `gorm:"column:id" json:"id"`
	PlateNumber string    `gorm:"column:plate_number" json:"plate_number"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}
