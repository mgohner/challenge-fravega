package route

import "github.com/google/uuid"

type CreateRoute struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	VehicleId   uuid.UUID `json:"vehicle_id"`
	DriverId    uuid.UUID `json:"driver_id"`
}
