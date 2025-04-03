package routePoint

import (
	"time"

	"github.com/google/uuid"
)

type RoutePoint struct {
	ID              uuid.UUID `gorm:"column:id"`
	PurchaseOrderID string    `gorm:"column:purchase_order_id"`
	RouteID         uuid.UUID `gorm:"column:route_id"`
	Status          string    `gorm:"column:status"`
	Latitude        float64   `gorm:"column:latitude"`
	Longitude       float64   `gorm:"column:longitude"`
	Address         string    `gorm:"column:address"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at"`
}

type RoutePointStatus string

const (
	RoutePointStatusPending   RoutePointStatus = "pending"
	RoutePointStatusInRoute   RoutePointStatus = "in_route"
	RoutePointStatusCompleted RoutePointStatus = "completed"
)

var RoutePointStatusList = map[RoutePointStatus]string{
	RoutePointStatusPending:   "pending",
	RoutePointStatusInRoute:   "in_route",
	RoutePointStatusCompleted: "completed",
}
