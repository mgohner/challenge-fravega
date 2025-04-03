package route

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) CreateRoute(route *Route) (*Route, error) {
	if route.ID == uuid.Nil {
		route.ID = uuid.New()
	}
	err := r.db.Create(route).Error
	return route, err
}

func (r *Repository) GetRoutes() ([]Route, error) {
	var routes []Route
	err := r.db.Preload("Vehicle").Preload("Driver").Preload("RoutePoints").Find(&routes).Error
	return routes, err
}

func (r *Repository) GetRoute(id string) (*Route, error) {
	var route Route
	err := r.db.Preload("Vehicle").Preload("Driver").Preload("RoutePoints").First(&route, "id = ?", id).Error
	return &route, err
}

// static functions

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
