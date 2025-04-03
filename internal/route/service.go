package route

type Service interface {
	GetRoutes() ([]Route, error)
	GetRoute(id string) (*Route, error)
	CreateRoute(newRoute *CreateRoute) (*Route, error)
}

type service struct {
	repository *Repository
}

func (s *service) GetRoutes() ([]Route, error) {
	return s.repository.GetRoutes()
}

func (s *service) GetRoute(id string) (*Route, error) {
	return s.repository.GetRoute(id)
}

func (s *service) CreateRoute(newRoute *CreateRoute) (*Route, error) {
	route, err := s.repository.CreateRoute(&Route{
		Name:        newRoute.Name,
		Description: newRoute.Description,
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   newRoute.VehicleId,
		DriverID:    newRoute.DriverId,
	})
	if err != nil {
		return nil, err
	}

	return s.repository.GetRoute(route.ID.String())
}

// static functions

func NewService(repository *Repository) *service {
	return &service{repository: repository}
}
