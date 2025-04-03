package carDriver

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateDriver(driver *Driver) (*Driver, error) {
	if driver.ID == uuid.Nil {
		driver.ID = uuid.New()
	}
	return driver, r.db.Create(driver).Error
}

func (r *Repository) GetDriver(id uuid.UUID) (*Driver, error) {
	var driver Driver
	return &driver, r.db.First(&driver, "id = ?", id).Error
}

func (r *Repository) GetDrivers() ([]Driver, error) {
	var drivers []Driver
	return drivers, r.db.Find(&drivers).Error
}

func (r *Repository) UpdateDriver(driver *Driver) (*Driver, error) {
	return driver, r.db.Save(driver).Error
}

// static functions

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
