package routePoint

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateRoutePoint(routePoint *RoutePoint) (*RoutePoint, error) {
	if routePoint.ID == uuid.Nil {
		routePoint.ID = uuid.New()
	}
	err := r.db.Create(routePoint).Error
	return routePoint, err
}

func (r *Repository) GetRoutePoints() ([]RoutePoint, error) {
	var routePoints []RoutePoint
	err := r.db.Find(&routePoints).Error
	return routePoints, err
}

func (r *Repository) GetRoutePoint(id string) (*RoutePoint, error) {
	var routePoint RoutePoint
	err := r.db.First(&routePoint, "id = ?", id).Error
	return &routePoint, err
}

// static functions

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
