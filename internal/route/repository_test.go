package route

import (
	carDriver "challenge-fravega/internal/car-driver"
	routePoint "challenge-fravega/internal/route-point"
	"challenge-fravega/internal/vehicle"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type RepositoryTestSuite struct {
	suite.Suite
	db         *gorm.DB
	repository *Repository
}

func (suite *RepositoryTestSuite) SetupTest() {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		suite.T().Fatal(err)
	}

	// Migrate all required tables
	err = db.AutoMigrate(
		&Route{},
		&vehicle.Vehicle{},
		&carDriver.Driver{},
		&routePoint.RoutePoint{},
	)
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.repository = NewRepository(db)
}

func (suite *RepositoryTestSuite) TestCreateRoute() {
	// Arrange
	// Create required vehicle and driver first
	vehicleID := uuid.New()
	vehicle := &vehicle.Vehicle{
		ID:          vehicleID,
		PlateNumber: "XYZ789",
	}

	driverID := uuid.New()
	driver := &carDriver.Driver{
		ID:             driverID,
		Name:           "John Doe",
		PhoneNumber:    "1234567890",
		Email:          "john@example.com",
		Address:        "123 Test St",
		Identification: "ID12345",
		LicenseNumber:  "LIC12345",
	}

	suite.db.Create(vehicle)
	suite.db.Create(driver)

	route := &Route{
		Name:        "Test Route",
		Description: "Test Route Description",
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	// Act
	result, err := suite.repository.CreateRoute(route)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), uuid.Nil, result.ID)
	assert.Equal(suite.T(), route.Name, result.Name)
	assert.Equal(suite.T(), route.Description, result.Description)
	assert.Equal(suite.T(), route.Status, result.Status)
	assert.Equal(suite.T(), vehicleID, result.VehicleID)
	assert.Equal(suite.T(), driverID, result.DriverID)
	assert.False(suite.T(), result.CreatedAt.IsZero())
}

func (suite *RepositoryTestSuite) TestGetRoute() {
	// Arrange
	// Create required vehicle and driver first
	vehicleID := uuid.New()
	vehicle := &vehicle.Vehicle{
		ID:          vehicleID,
		PlateNumber: "ABC123",
	}

	driverID := uuid.New()
	driver := &carDriver.Driver{
		ID:             driverID,
		Name:           "Jane Doe",
		PhoneNumber:    "0987654321",
		Email:          "jane@example.com",
		Address:        "456 Test St",
		Identification: "ID54321",
		LicenseNumber:  "LIC54321",
	}

	suite.db.Create(vehicle)
	suite.db.Create(driver)

	routeID := uuid.New()
	route := &Route{
		ID:          routeID,
		Name:        "Test Route",
		Description: "Test Route Description",
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.db.Create(route)

	// Act
	result, err := suite.repository.GetRoute(routeID.String())

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), routeID, result.ID)
	assert.Equal(suite.T(), route.Name, result.Name)
	assert.Equal(suite.T(), route.Description, result.Description)
	assert.Equal(suite.T(), vehicleID, result.VehicleID)
	assert.Equal(suite.T(), driverID, result.DriverID)
}

func (suite *RepositoryTestSuite) TestGetRoutes() {
	// Arrange
	// Clear any existing routes
	suite.db.Exec("DELETE FROM routes")

	// Create required vehicle and driver first
	vehicleID := uuid.New()
	vehicle := &vehicle.Vehicle{
		ID:          vehicleID,
		PlateNumber: "MNO456",
	}

	driverID := uuid.New()
	driver := &carDriver.Driver{
		ID:             driverID,
		Name:           "Test Driver",
		PhoneNumber:    "5555555555",
		Email:          "test@example.com",
		Address:        "789 Test St",
		Identification: "ID99999",
		LicenseNumber:  "LIC99999",
	}

	suite.db.Create(vehicle)
	suite.db.Create(driver)

	route1 := &Route{
		ID:          uuid.New(),
		Name:        "Route 1",
		Description: "Route 1 Description",
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	route2 := &Route{
		ID:          uuid.New(),
		Name:        "Route 2",
		Description: "Route 2 Description",
		Status:      RouteStatusList[RouteStatusStarted],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	suite.db.Create(route1)
	suite.db.Create(route2)

	// Act
	results, err := suite.repository.GetRoutes()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 2)
}

func (suite *RepositoryTestSuite) TestGetRouteWithRelations() {
	// Arrange
	// Create required vehicle and driver first
	vehicleID := uuid.New()
	vehicle := &vehicle.Vehicle{
		ID:          vehicleID,
		PlateNumber: "XYZ999",
	}

	driverID := uuid.New()
	driver := &carDriver.Driver{
		ID:             driverID,
		Name:           "Test Driver",
		PhoneNumber:    "1112223333",
		Email:          "test@example.com",
		Address:        "999 Test St",
		Identification: "ID8888",
		LicenseNumber:  "LIC8888",
	}

	suite.db.Create(vehicle)
	suite.db.Create(driver)

	routeID := uuid.New()
	route := &Route{
		ID:          routeID,
		Name:        "Route With Relations",
		Description: "Route With Relations Description",
		Status:      RouteStatusList[RouteStatusPending],
		VehicleID:   vehicleID,
		DriverID:    driverID,
	}

	suite.db.Create(route)

	// Create a route point for this route
	routePoint1 := &routePoint.RoutePoint{
		ID:              uuid.New(),
		PurchaseOrderID: "PO12345",
		RouteID:         routeID,
		Status:          routePoint.RoutePointStatusList[routePoint.RoutePointStatusPending],
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test Point",
	}

	suite.db.Create(routePoint1)

	// Act
	result, err := suite.repository.GetRoute(routeID.String())

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), routeID, result.ID)
	assert.Equal(suite.T(), route.Name, result.Name)

	// Check vehicle relation
	assert.Equal(suite.T(), vehicleID, result.Vehicle.ID)
	assert.Equal(suite.T(), vehicle.PlateNumber, result.Vehicle.PlateNumber)

	// Check driver relation
	assert.Equal(suite.T(), driverID, result.Driver.ID)
	assert.Equal(suite.T(), driver.Name, result.Driver.Name)

	// Check route points relation
	assert.GreaterOrEqual(suite.T(), len(result.RoutePoints), 1)

	// Find the route point we created
	var foundRoutePoint *routePoint.RoutePoint
	for _, rp := range result.RoutePoints {
		if rp.ID == routePoint1.ID {
			foundRoutePoint = &rp
			break
		}
	}

	assert.NotNil(suite.T(), foundRoutePoint)
	assert.Equal(suite.T(), routePoint1.PurchaseOrderID, foundRoutePoint.PurchaseOrderID)
	assert.Equal(suite.T(), routePoint1.Address, foundRoutePoint.Address)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
