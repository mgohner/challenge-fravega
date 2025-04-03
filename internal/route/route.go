package route

import (
	carDriver "challenge-fravega/internal/car-driver"
	routePoint "challenge-fravega/internal/route-point"
	"challenge-fravega/internal/vehicle"
	"time"

	"github.com/google/uuid"
)

type Route struct {
	ID          uuid.UUID               `gorm:"column:id"`
	Name        string                  `gorm:"column:name"`
	Description string                  `gorm:"column:description"`
	Status      string                  `gorm:"column:status"`
	CreatedAt   time.Time               `gorm:"column:created_at"`
	UpdatedAt   time.Time               `gorm:"column:updated_at"`
	VehicleID   uuid.UUID               `gorm:"column:vehicle_id"`
	Vehicle     vehicle.Vehicle         `gorm:"foreignKey:ID;references:VehicleID"`
	DriverID    uuid.UUID               `gorm:"column:driver_id"`
	Driver      carDriver.Driver        `gorm:"foreignKey:ID;references:DriverID"`
	RoutePoints []routePoint.RoutePoint `gorm:"foreignKey:RouteID"`
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
