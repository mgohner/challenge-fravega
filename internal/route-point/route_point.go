package routePoint

import (
	"time"

	"github.com/google/uuid"
)

type RoutePoint struct {
	ID              uuid.UUID `gorm:"column:id" json:"id"`
	PurchaseOrderID string    `gorm:"column:purchase_order_id" json:"purchase_order_id"`
	RouteID         uuid.UUID `gorm:"column:route_id" json:"route_id"`
	Status          string    `gorm:"column:status" json:"status"`
	Latitude        float64   `gorm:"column:latitude" json:"latitude"`
	Longitude       float64   `gorm:"column:longitude" json:"longitude"`
	Address         string    `gorm:"column:address" json:"address"`
	CreatedAt       time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time `gorm:"column:updated_at" json:"updated_at"`
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
