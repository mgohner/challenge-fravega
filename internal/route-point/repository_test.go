package routePoint

import (
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

	err = db.AutoMigrate(&RoutePoint{})
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.repository = NewRepository(db)
}

func (suite *RepositoryTestSuite) TestCreateRoutePoint() {
	// Arrange
	routeID := uuid.New()
	routePoint := &RoutePoint{
		PurchaseOrderID: "PO12345",
		RouteID:         routeID,
		Status:          RoutePointStatusList[RoutePointStatusPending],
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
	}

	// Act
	result, err := suite.repository.CreateRoutePoint(routePoint)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), uuid.Nil, result.ID)
	assert.Equal(suite.T(), routePoint.PurchaseOrderID, result.PurchaseOrderID)
	assert.Equal(suite.T(), routePoint.RouteID, result.RouteID)
	assert.Equal(suite.T(), routePoint.Status, result.Status)
	assert.Equal(suite.T(), routePoint.Latitude, result.Latitude)
	assert.Equal(suite.T(), routePoint.Longitude, result.Longitude)
	assert.Equal(suite.T(), routePoint.Address, result.Address)
	assert.False(suite.T(), result.CreatedAt.IsZero())
}

func (suite *RepositoryTestSuite) TestGetRoutePoint() {
	// Arrange
	routePointID := uuid.New()
	routeID := uuid.New()
	routePoint := &RoutePoint{
		ID:              routePointID,
		PurchaseOrderID: "PO12345",
		RouteID:         routeID,
		Status:          RoutePointStatusList[RoutePointStatusPending],
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	suite.db.Create(routePoint)

	// Act
	result, err := suite.repository.GetRoutePoint(routePointID.String())

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), routePointID, result.ID)
	assert.Equal(suite.T(), routePoint.PurchaseOrderID, result.PurchaseOrderID)
	assert.Equal(suite.T(), routePoint.RouteID, result.RouteID)
	assert.Equal(suite.T(), routePoint.Status, result.Status)
	assert.Equal(suite.T(), routePoint.Latitude, result.Latitude)
	assert.Equal(suite.T(), routePoint.Longitude, result.Longitude)
	assert.Equal(suite.T(), routePoint.Address, result.Address)
}

func (suite *RepositoryTestSuite) TestGetRoutePoints() {
	// Arrange
	// Clear any existing routePoints
	suite.db.Exec("DELETE FROM route_points")

	routeID := uuid.New()
	routePoint1 := &RoutePoint{
		ID:              uuid.New(),
		PurchaseOrderID: "PO12345",
		RouteID:         routeID,
		Status:          RoutePointStatusList[RoutePointStatusPending],
		Latitude:        37.7749,
		Longitude:       -122.4194,
		Address:         "123 Test St",
	}

	routePoint2 := &RoutePoint{
		ID:              uuid.New(),
		PurchaseOrderID: "PO67890",
		RouteID:         routeID,
		Status:          RoutePointStatusList[RoutePointStatusInRoute],
		Latitude:        34.0522,
		Longitude:       -118.2437,
		Address:         "456 Test Ave",
	}

	suite.db.Create(routePoint1)
	suite.db.Create(routePoint2)

	// Act
	results, err := suite.repository.GetRoutePoints()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 2)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
