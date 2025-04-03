package vehicle

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a repository interface that our mock can implement
type RepositoryInterface interface {
	CreateVehicle(vehicle *Vehicle) (*Vehicle, error)
	GetVehicle(id uuid.UUID) (*Vehicle, error)
	GetVehicles() ([]Vehicle, error)
	UpdateVehicle(vehicle *Vehicle) (*Vehicle, error)
}

// Define a mock repository for testing the service
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	args := m.Called(vehicle)
	return args.Get(0).(*Vehicle), args.Error(1)
}

func (m *MockRepository) GetVehicle(id uuid.UUID) (*Vehicle, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Vehicle), args.Error(1)
}

func (m *MockRepository) GetVehicles() ([]Vehicle, error) {
	args := m.Called()
	return args.Get(0).([]Vehicle), args.Error(1)
}

func (m *MockRepository) UpdateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	args := m.Called(vehicle)
	return args.Get(0).(*Vehicle), args.Error(1)
}

// Create a custom service for testing
type testService struct {
	repo RepositoryInterface
}

func (s *testService) CreateVehicle(vehicle *Vehicle) (*Vehicle, error) {
	return s.repo.CreateVehicle(vehicle)
}

func (s *testService) GetVehicle(id uuid.UUID) (*Vehicle, error) {
	return s.repo.GetVehicle(id)
}

func (s *testService) GetVehicles() ([]Vehicle, error) {
	return s.repo.GetVehicles()
}

func createTestService(mockRepo *MockRepository) Service {
	return &testService{repo: mockRepo}
}

func TestCreateVehicle(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	vehicle := &Vehicle{
		PlateNumber: "ABC123",
	}

	expectedID := uuid.New()
	expectedVehicle := &Vehicle{
		ID:          expectedID,
		PlateNumber: vehicle.PlateNumber,
	}

	mockRepo.On("CreateVehicle", vehicle).Return(expectedVehicle, nil)

	// Act
	result, err := service.CreateVehicle(vehicle)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedID, result.ID)
	assert.Equal(t, vehicle.PlateNumber, result.PlateNumber)
	mockRepo.AssertExpectations(t)
}

func TestGetVehicle(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	vehicleID := uuid.New()
	expectedVehicle := &Vehicle{
		ID:          vehicleID,
		PlateNumber: "ABC123",
	}

	mockRepo.On("GetVehicle", vehicleID).Return(expectedVehicle, nil)

	// Act
	result, err := service.GetVehicle(vehicleID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, vehicleID, result.ID)
	assert.Equal(t, expectedVehicle.PlateNumber, result.PlateNumber)
	mockRepo.AssertExpectations(t)
}

func TestGetVehicleNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	vehicleID := uuid.New()
	expectedError := errors.New("vehicle not found")

	mockRepo.On("GetVehicle", vehicleID).Return((*Vehicle)(nil), expectedError)

	// Act
	_, err := service.GetVehicle(vehicleID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetVehicles(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	expectedVehicles := []Vehicle{
		{
			ID:          uuid.New(),
			PlateNumber: "ABC123",
		},
		{
			ID:          uuid.New(),
			PlateNumber: "XYZ789",
		},
	}

	mockRepo.On("GetVehicles").Return(expectedVehicles, nil)

	// Act
	results, err := service.GetVehicles()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, expectedVehicles[0].PlateNumber, results[0].PlateNumber)
	assert.Equal(t, expectedVehicles[1].PlateNumber, results[1].PlateNumber)
	mockRepo.AssertExpectations(t)
}
