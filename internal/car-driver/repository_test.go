package carDriver

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

	err = db.AutoMigrate(&Driver{})
	if err != nil {
		suite.T().Fatal(err)
	}

	suite.db = db
	suite.repository = NewRepository(db)
}

func (suite *RepositoryTestSuite) TestCreateDriver() {
	// Arrange
	driver := &Driver{
		Name:           "John Doe",
		PhoneNumber:    "1234567890",
		Email:          "john@example.com",
		Address:        "123 Main St",
		Identification: "ID12345",
		LicenseNumber:  "LIC98765",
	}

	// Act
	result, err := suite.repository.CreateDriver(driver)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotEqual(suite.T(), uuid.Nil, result.ID)
	assert.Equal(suite.T(), driver.Name, result.Name)
	assert.Equal(suite.T(), driver.PhoneNumber, result.PhoneNumber)
	assert.Equal(suite.T(), driver.Email, result.Email)
	assert.Equal(suite.T(), driver.Address, result.Address)
	assert.Equal(suite.T(), driver.Identification, result.Identification)
	assert.Equal(suite.T(), driver.LicenseNumber, result.LicenseNumber)
	assert.False(suite.T(), result.CreatedAt.IsZero())
}

func (suite *RepositoryTestSuite) TestGetDriver() {
	// Arrange
	driverID := uuid.New()
	driver := &Driver{
		ID:             driverID,
		Name:           "Jane Doe",
		PhoneNumber:    "0987654321",
		Email:          "jane@example.com",
		Address:        "456 Oak St",
		Identification: "ID54321",
		LicenseNumber:  "LIC56789",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	suite.db.Create(driver)

	// Act
	result, err := suite.repository.GetDriver(driverID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), driverID, result.ID)
	assert.Equal(suite.T(), driver.Name, result.Name)
	assert.Equal(suite.T(), driver.PhoneNumber, result.PhoneNumber)
	assert.Equal(suite.T(), driver.Email, result.Email)
}

func (suite *RepositoryTestSuite) TestGetDrivers() {
	// Arrange
	// Clear any existing drivers
	suite.db.Exec("DELETE FROM drivers")

	driver1 := &Driver{
		ID:             uuid.New(),
		Name:           "Driver One",
		PhoneNumber:    "1111111111",
		Email:          "one@example.com",
		Identification: "ID1111",
		LicenseNumber:  "LIC1111",
	}

	driver2 := &Driver{
		ID:             uuid.New(),
		Name:           "Driver Two",
		PhoneNumber:    "2222222222",
		Email:          "two@example.com",
		Identification: "ID2222",
		LicenseNumber:  "LIC2222",
	}

	suite.db.Create(driver1)
	suite.db.Create(driver2)

	// Act
	results, err := suite.repository.GetDrivers()

	// Assert
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), results, 2)
}

func (suite *RepositoryTestSuite) TestUpdateDriver() {
	// Arrange
	driverID := uuid.New()
	driver := &Driver{
		ID:             driverID,
		Name:           "Update Test",
		PhoneNumber:    "5555555555",
		Email:          "update@example.com",
		Address:        "789 Update St",
		Identification: "ID5555",
		LicenseNumber:  "LIC5555",
	}

	suite.db.Create(driver)

	// Update values
	driver.Name = "Updated Name"
	driver.PhoneNumber = "9999999999"

	// Act
	result, err := suite.repository.UpdateDriver(driver)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Updated Name", result.Name)
	assert.Equal(suite.T(), "9999999999", result.PhoneNumber)

	// Verify in database
	var updatedDriver Driver
	suite.db.First(&updatedDriver, "id = ?", driverID)
	assert.Equal(suite.T(), "Updated Name", updatedDriver.Name)
	assert.Equal(suite.T(), "9999999999", updatedDriver.PhoneNumber)
}

func TestRepositorySuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
