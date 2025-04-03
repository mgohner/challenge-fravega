package routePoint

type Service interface {
	GetRoutePoints() ([]RoutePoint, error)
	GetRoutePoint(id string) (*RoutePoint, error)
	CreateRoutePoint(addPurchaseOrder *AddPurchaseOrder) (*RoutePoint, error)
}

type service struct {
	repository *Repository
}

func (s *service) GetRoutePoints() ([]RoutePoint, error) {
	return s.repository.GetRoutePoints()
}

func (s *service) GetRoutePoint(id string) (*RoutePoint, error) {
	return s.repository.GetRoutePoint(id)
}

func (s *service) CreateRoutePoint(addPurchaseOrder *AddPurchaseOrder) (*RoutePoint, error) {
	routePoint := &RoutePoint{
		RouteID:         addPurchaseOrder.RouteID,
		PurchaseOrderID: addPurchaseOrder.PurchaseOrderID,
		Latitude:        addPurchaseOrder.Latitude,
		Longitude:       addPurchaseOrder.Longitude,
		Address:         addPurchaseOrder.Address,
		Status:          RoutePointStatusList[RoutePointStatusPending],
	}
	return s.repository.CreateRoutePoint(routePoint)
}

// static functions

func NewService(repository *Repository) *service {
	return &service{repository: repository}
}
