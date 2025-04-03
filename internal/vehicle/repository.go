package vehicle

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	if vehicle.ID == uuid.Nil {
		vehicle.ID = uuid.New()
	}
	return vehicle, r.db.Create(vehicle).Error
}

func (r *Repository) GetVehicle(id uuid.UUID) (*Vehicle, error) {
	var vehicle Vehicle
	return &vehicle, r.db.First(&vehicle, "id = ?", id).Error
}

func (r *Repository) GetVehicles() ([]Vehicle, error) {
	var vehicles []Vehicle
	return vehicles, r.db.Find(&vehicles).Error
}

func (r *Repository) UpdateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	return vehicle, r.db.Save(vehicle).Error
}

// static functions

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
