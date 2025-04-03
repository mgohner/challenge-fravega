package route

import (
	carDriver "challenge-fravega/internal/car-driver"
	routePoint "challenge-fravega/internal/route-point"
	"challenge-fravega/internal/vehicle"
	"time"

	"github.com/google/uuid"
)

type Route struct {
	ID          uuid.UUID               `gorm:"column:id" json:"id"`
	Name        string                  `gorm:"column:name" json:"name"`
	Description string                  `gorm:"column:description" json:"description"`
	Status      string                  `gorm:"column:status" json:"status"`
	CreatedAt   time.Time               `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time               `gorm:"column:updated_at" json:"updated_at"`
	VehicleID   uuid.UUID               `gorm:"column:vehicle_id" json:"vehicle_id"`
	Vehicle     vehicle.Vehicle         `gorm:"foreignKey:ID;references:VehicleID" json:"vehicle"`
	DriverID    uuid.UUID               `gorm:"column:driver_id" json:"driver_id"`
	Driver      carDriver.Driver        `gorm:"foreignKey:ID;references:DriverID" json:"driver"`
	RoutePoints []routePoint.RoutePoint `gorm:"foreignKey:RouteID" json:"route_points"`
}

type RouteStatus string

const (
	RouteStatusPending   RouteStatus = "pending"
	RouteStatusStarted   RouteStatus = "started"
	RouteStatusCompleted RouteStatus = "completed"
)

var RouteStatusList = map[RouteStatus]string{
	RouteStatusPending:   "pending",
	RouteStatusStarted:   "started",
	RouteStatusCompleted: "completed",
}
