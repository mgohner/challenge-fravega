package routePoint

import "github.com/google/uuid"

type AddPurchaseOrder struct {
	RouteID         uuid.UUID `json:"route_id"`
	PurchaseOrderID string    `json:"purchase_order_id"`
	Latitude        float64   `json:"latitude"`
	Longitude       float64   `json:"longitude"`
	Address         string    `json:"address"`
}
