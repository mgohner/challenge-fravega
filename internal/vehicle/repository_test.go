package vehicle

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

	err = db.AutoMigrate(&Vehicle{})
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.repository = NewRepository(db)
}

func (suite *RepositoryTestSuite) TestCreateVehicle() {
	// Arrange
	vehicle := &Vehicle{
		PlateNumber: "XYZ789",
	}

	// Act
	result, err := suite.repository.CreateVehicle(vehicle)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), uuid.Nil, result.ID)
	assert.Equal(suite.T(), vehicle.PlateNumber, result.PlateNumber)
	assert.False(suite.T(), result.CreatedAt.IsZero())
}

func (suite *RepositoryTestSuite) TestGetVehicle() {
	// Arrange
	vehicleID := uuid.New()
	vehicle := &Vehicle{
		ID:          vehicleID,
		PlateNumber: "ABC123",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	suite.db.Create(vehicle)

	// Act
	result, err := suite.repository.GetVehicle(vehicleID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), vehicleID, result.ID)
	assert.Equal(suite.T(), vehicle.PlateNumber, result.PlateNumber)
}

func (suite *RepositoryTestSuite) TestGetVehicles() {
	// Arrange
	// Clear any existing vehicles
	suite.db.Exec("DELETE FROM vehicles")

	vehicle1 := &Vehicle{
		ID:          uuid.New(),
		PlateNumber: "ABC123",
	}

	vehicle2 := &Vehicle{
		ID:          uuid.New(),
		PlateNumber: "XYZ789",
	}

	suite.db.Create(vehicle1)
	suite.db.Create(vehicle2)

	// Act
	results, err := suite.repository.GetVehicles()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 2)
}

func (suite *RepositoryTestSuite) TestUpdateVehicle() {
	// Arrange
	vehicleID := uuid.New()
	vehicle := &Vehicle{
		ID:          vehicleID,
		PlateNumber: "OLD123",
	}

	suite.db.Create(vehicle)

	// Update values
	vehicle.PlateNumber = "NEW456"

	// Act
	result, err := suite.repository.UpdateVehicle(vehicle)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "NEW456", result.PlateNumber)

	// Verify in database
	var updatedVehicle Vehicle
	suite.db.First(&updatedVehicle, "id = ?", vehicleID)
	assert.Equal(suite.T(), "NEW456", updatedVehicle.PlateNumber)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
