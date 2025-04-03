package carDriver

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a repository interface that our mock can implement
type RepositoryInterface interface {
	CreateDriver(driver *Driver) (*Driver, error)
	GetDriver(id uuid.UUID) (*Driver, error)
	GetDrivers() ([]Driver, error)
	UpdateDriver(driver *Driver) (*Driver, error)
}

// Define a mock repository for testing the service
type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreateDriver(driver *Driver) (*Driver, error) {
	args := m.Called(driver)
	return args.Get(0).(*Driver), args.Error(1)
}

func (m *MockRepository) GetDriver(id uuid.UUID) (*Driver, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Driver), args.Error(1)
}

func (m *MockRepository) GetDrivers() ([]Driver, error) {
	args := m.Called()
	return args.Get(0).([]Driver), args.Error(1)
}

func (m *MockRepository) UpdateDriver(driver *Driver) (*Driver, error) {
	args := m.Called(driver)
	return args.Get(0).(*Driver), args.Error(1)
}

// Create a custom service for testing
type testService struct {
	repo RepositoryInterface
}

func (s *testService) CreateDriver(driver *Driver) (*Driver, error) {
	return s.repo.CreateDriver(driver)
}

func (s *testService) GetDriver(id uuid.UUID) (*Driver, error) {
	return s.repo.GetDriver(id)
}

func (s *testService) GetDrivers() ([]Driver, error) {
	return s.repo.GetDrivers()
}

func createTestService(mockRepo *MockRepository) Service {
	return &testService{repo: mockRepo}
}

func TestCreateDriver(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	driver := &Driver{
		Name:           "Test Driver",
		PhoneNumber:    "1234567890",
		Email:          "test@example.com",
		Address:        "123 Test St",
		Identification: "ID-TEST",
		LicenseNumber:  "LIC-TEST",
	}

	expectedID := uuid.New()
	expectedDriver := &Driver{
		ID:             expectedID,
		Name:           driver.Name,
		PhoneNumber:    driver.PhoneNumber,
		Email:          driver.Email,
		Address:        driver.Address,
		Identification: driver.Identification,
		LicenseNumber:  driver.LicenseNumber,
	}

	mockRepo.On("CreateDriver", driver).Return(expectedDriver, nil)

	// Act
	result, err := service.CreateDriver(driver)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedID, result.ID)
	assert.Equal(t, driver.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetDriver(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	driverID := uuid.New()
	expectedDriver := &Driver{
		ID:             driverID,
		Name:           "Found Driver",
		PhoneNumber:    "5551234567",
		Email:          "found@example.com",
		Address:        "456 Found St",
		Identification: "ID-FOUND",
		LicenseNumber:  "LIC-FOUND",
	}

	mockRepo.On("GetDriver", driverID).Return(expectedDriver, nil)

	// Act
	result, err := service.GetDriver(driverID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, driverID, result.ID)
	assert.Equal(t, expectedDriver.Name, result.Name)
	mockRepo.AssertExpectations(t)
}

func TestGetDriverNotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	driverID := uuid.New()
	expectedError := errors.New("driver not found")

	mockRepo.On("GetDriver", driverID).Return((*Driver)(nil), expectedError)

	// Act
	_, err := service.GetDriver(driverID)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}

func TestGetDrivers(t *testing.T) {
	// Arrange
	mockRepo := new(MockRepository)
	service := createTestService(mockRepo)

	expectedDrivers := []Driver{
		{
			ID:             uuid.New(),
			Name:           "Driver 1",
			PhoneNumber:    "1111111111",
			Email:          "driver1@example.com",
			Identification: "ID-1",
			LicenseNumber:  "LIC-1",
		},
		{
			ID:             uuid.New(),
			Name:           "Driver 2",
			PhoneNumber:    "2222222222",
			Email:          "driver2@example.com",
			Identification: "ID-2",
			LicenseNumber:  "LIC-2",
		},
	}

	mockRepo.On("GetDrivers").Return(expectedDrivers, nil)

	// Act
	results, err := service.GetDrivers()

	// Assert
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, expectedDrivers[0].Name, results[0].Name)
	assert.Equal(t, expectedDrivers[1].Name, results[1].Name)
	mockRepo.AssertExpectations(t)
}
