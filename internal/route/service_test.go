package route

import (
	carDriver "challenge-fravega/internal/car-driver"
	"challenge-fravega/internal/vehicle"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a repository interface that our mock can implement
type RepositoryInterface interface {
	CreateRoute(route *Route) (*Route, error)
	GetRoute(id string) (*Route, error)
	GetRoutes() ([]Route, error)
}

// Define a mock repository for testing the service
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateRoute(route *Route) (*Route, error) {
	args := m.Called(route)
	return args.Get(0).(*Route), args.Error(1)
}

func (m *MockRepository) GetRoute(id string) (*Route, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Route), args.Error(1)
}

func (m *MockRepository) GetRoutes() ([]Route, error) {
	args := m.Called()
	return args.Get(0).([]Route), args.Error(1)
}

// Create a custom service for testing
type testService struct {
	repo RepositoryInterface
}

func (s *testService) CreateRoute(newRoute *CreateRoute) (*Route, error) {
	route := &Route{
		Name:        newRoute.Name,
		Description: newRoute.Description,
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   newRoute.VehicleId,
		DriverID:    newRoute.DriverId,
	}

	createdRoute, err := s.repo.CreateRoute(route)
	if err != nil {
		return nil, err
	}

	return s.repo.GetRoute(createdRoute.ID.String())
}

func (s *testService) GetRoute(id string) (*Route, error) {
	return s.repo.GetRoute(id)
}

func (s *testService) GetRoutes() ([]Route, error) {
	return s.repo.GetRoutes()
}

func createTestService(mockRepo *MockRepository) Service {
	return &testService{repo: mockRepo}
}

func TestCreateRoute(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	vehicleID := uuid.New()
	driverID := uuid.New()
	routeID := uuid.New()

	createRequest := &CreateRoute{
		Name:        "New Route",
		Description: "New Route Description",
		VehicleId:   vehicleID,
		DriverId:    driverID,
	}

	// The route that will be created and passed to the repository
	expectedRouteCreate := &Route{
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	// The route that will be returned after creation
	createdRoute := &Route{
		ID:          routeID,
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	// The route that will be returned by GetRoute
	fullRoute := &Route{
		ID:          routeID,
		Name:        createRequest.Name,
		Description: createRequest.Description,
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
		Vehicle:     vehicle.Vehicle{ID: vehicleID, PlateNumber: "ABC123"},
		Driver:      carDriver.Driver{ID: driverID, Name: "John Doe"},
	}

	// Mock the repository calls
	mockRepo.On("CreateRoute", mock.MatchedBy(func(r *Route) bool {
		return r.Name == expectedRouteCreate.Name &&
			r.Description == expectedRouteCreate.Description &&
			r.Status == expectedRouteCreate.Status &&
			r.VehicleID == expectedRouteCreate.VehicleID &&
			r.DriverID == expectedRouteCreate.DriverID
	})).Return(createdRoute, nil)

	mockRepo.On("GetRoute", routeID.String()).Return(fullRoute, nil)

	// Act
	result, err := service.CreateRoute(createRequest)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, routeID, result.ID)
	assert.Equal(t, createRequest.Name, result.Name)
	assert.Equal(t, createRequest.Description, result.Description)
	assert.Equal(t, RouteStatusList[RouteStatusPending], result.Status)
	assert.Equal(t, vehicleID, result.VehicleID)
	assert.Equal(t, driverID, result.DriverID)
	mockRepo.AssertExpectations(t)
}

func TestGetRoute(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routeID := uuid.New()
	expectedRoute := &Route{
		ID:          routeID,
		Name:        "Test Route",
		Description: "Test Route Description",
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   uuid.New(),
		DriverID:    uuid.New(),
	}

	mockRepo.On("GetRoute", routeID.String()).Return(expectedRoute, nil)

	// Act
	result, err := service.GetRoute(routeID.String())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, routeID, result.ID)
	assert.Equal(t, expectedRoute.Name, result.Name)
	assert.Equal(t, expectedRoute.Description, result.Description)
	mockRepo.AssertExpectations(t)
}

func TestGetRouteNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routeID := uuid.New()
	expectedError := errors.New("route not found")

	mockRepo.On("GetRoute", routeID.String()).Return((*Route)(nil), expectedError)

	// Act
	_, err := service.GetRoute(routeID.String())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetRoutes(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	expectedRoutes := []Route{
		{
			ID:          uuid.New(),
			Name:        "Route 1",
			Description: "Route 1 Description",
			Status:      RouteStatusList[RouteStatusPending],
			VehicleID:   uuid.New(),
			DriverID:    uuid.New(),
		},
		{
			ID:          uuid.New(),
			Name:        "Route 2",
			Description: "Route 2 Description",
			Status:      RouteStatusList[RouteStatusStarted],
			VehicleID:   uuid.New(),
			DriverID:    uuid.New(),
		},
	}

	mockRepo.On("GetRoutes").Return(expectedRoutes, nil)

	// Act
	results, err := service.GetRoutes()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, expectedRoutes[0].Name, results[0].Name)
	assert.Equal(t, expectedRoutes[1].Name, results[1].Name)
	mockRepo.AssertExpectations(t)
}
