package routePoint

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a repository interface that our mock can implement
type RepositoryInterface interface {
	CreateRoutePoint(routePoint *RoutePoint) (*RoutePoint, error)
	GetRoutePoint(id string) (*RoutePoint, error)
	GetRoutePoints() ([]RoutePoint, error)
}

// Define a mock repository for testing the service
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateRoutePoint(routePoint *RoutePoint) (*RoutePoint, error) {
	args := m.Called(routePoint)
	return args.Get(0).(*RoutePoint), args.Error(1)
}

func (m *MockRepository) GetRoutePoint(id string) (*RoutePoint, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*RoutePoint), args.Error(1)
}

func (m *MockRepository) GetRoutePoints() ([]RoutePoint, error) {
	args := m.Called()
	return args.Get(0).([]RoutePoint), args.Error(1)
}

// Create a custom service for testing
type testService struct {
	repo RepositoryInterface
}

func (s *testService) CreateRoutePoint(addPurchaseOrder *AddPurchaseOrder) (*RoutePoint, error) {
	routePoint := &RoutePoint{
		RouteID:         addPurchaseOrder.RouteID,
		PurchaseOrderID: addPurchaseOrder.PurchaseOrderID,
		Latitude:        addPurchaseOrder.Latitude,
		Longitude:       addPurchaseOrder.Longitude,
		Address:         addPurchaseOrder.Address,
		Status:          RoutePointStatusList[RoutePointStatusPending],
	}
	return s.repo.CreateRoutePoint(routePoint)
}

func (s *testService) GetRoutePoint(id string) (*RoutePoint, error) {
	return s.repo.GetRoutePoint(id)
}

func (s *testService) GetRoutePoints() ([]RoutePoint, error) {
	return s.repo.GetRoutePoints()
}

func createTestService(mockRepo *MockRepository) Service {
	return &testService{repo: mockRepo}
}

func TestCreateRoutePoint(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routeID := uuid.New()
	addPurchaseOrder := &AddPurchaseOrder{
		RouteID:         routeID,
		PurchaseOrderID: "PO12345",
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
	}

	// The route point that will be created in the service
	expectedRoutePoint := &RoutePoint{
		RouteID:         routeID,
		PurchaseOrderID: "PO12345",
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
		Status:          RoutePointStatusList[RoutePointStatusPending],
	}

	// The route point that will be returned by the repository
	createdRoutePoint := &RoutePoint{
		ID:              uuid.New(),
		RouteID:         routeID,
		PurchaseOrderID: "PO12345",
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
		Status:          RoutePointStatusList[RoutePointStatusPending],
	}

	// Mock the repository call
	mockRepo.On("CreateRoutePoint", mock.MatchedBy(func(rp *RoutePoint) bool {
		return rp.RouteID == expectedRoutePoint.RouteID &&
			rp.PurchaseOrderID == expectedRoutePoint.PurchaseOrderID &&
			rp.Latitude == expectedRoutePoint.Latitude &&
			rp.Longitude == expectedRoutePoint.Longitude &&
			rp.Address == expectedRoutePoint.Address &&
			rp.Status == expectedRoutePoint.Status
	})).Return(createdRoutePoint, nil)

	// Act
	result, err := service.CreateRoutePoint(addPurchaseOrder)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createdRoutePoint.ID, result.ID)
	assert.Equal(t, addPurchaseOrder.RouteID, result.RouteID)
	assert.Equal(t, addPurchaseOrder.PurchaseOrderID, result.PurchaseOrderID)
	assert.Equal(t, addPurchaseOrder.Latitude, result.Latitude)
	assert.Equal(t, addPurchaseOrder.Longitude, result.Longitude)
	assert.Equal(t, addPurchaseOrder.Address, result.Address)
	assert.Equal(t, RoutePointStatusList[RoutePointStatusPending], result.Status)
	mockRepo.AssertExpectations(t)
}

func TestGetRoutePoint(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routePointID := uuid.New()
	routeID := uuid.New()
	expectedRoutePoint := &RoutePoint{
		ID:              routePointID,
		RouteID:         routeID,
		PurchaseOrderID: "PO12345",
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
		Status:          RoutePointStatusList[RoutePointStatusPending],
	}

	mockRepo.On("GetRoutePoint", routePointID.String()).Return(expectedRoutePoint, nil)

	// Act
	result, err := service.GetRoutePoint(routePointID.String())

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, routePointID, result.ID)
	assert.Equal(t, routeID, result.RouteID)
	assert.Equal(t, "PO12345", result.PurchaseOrderID)
	assert.Equal(t, 37.7749, result.Latitude)
	assert.Equal(t, -122.4194, result.Longitude)
	assert.Equal(t, "123 Test St", result.Address)
	assert.Equal(t, RoutePointStatusList[RoutePointStatusPending], result.Status)
	mockRepo.AssertExpectations(t)
}

func TestGetRoutePointNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routePointID := uuid.New()
	expectedError := errors.New("route point not found")

	mockRepo.On("GetRoutePoint", routePointID.String()).Return((*RoutePoint)(nil), expectedError)

	// Act
	_, err := service.GetRoutePoint(routePointID.String())

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetRoutePoints(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	routeID := uuid.New()
	expectedRoutePoints := []RoutePoint{
		{
			ID:              uuid.New(),
			RouteID:         routeID,
			PurchaseOrderID: "PO12345",
			Latitude:        37.7749,
			Longitude:       -122.4194,
			Address:         "123 Test St",
			Status:          RoutePointStatusList[RoutePointStatusPending],
		},
		{
			ID:              uuid.New(),
			RouteID:         routeID,
			PurchaseOrderID: "PO67890",
			Latitude:        34.0522,
			Longitude:       -118.2437,
			Address:         "456 Test Ave",
			Status:          RoutePointStatusList[RoutePointStatusInRoute],
		},
	}

	mockRepo.On("GetRoutePoints").Return(expectedRoutePoints, nil)

	// Act
	results, err := service.GetRoutePoints()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, expectedRoutePoints[0].PurchaseOrderID, results[0].PurchaseOrderID)
	assert.Equal(t, expectedRoutePoints[1].PurchaseOrderID, results[1].PurchaseOrderID)
	mockRepo.AssertExpectations(t)
}
