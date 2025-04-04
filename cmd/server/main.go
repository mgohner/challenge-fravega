package main

import (
	"challenge-fravega/cmd/server/handlers"
	carDriver "challenge-fravega/internal/car-driver"
	"challenge-fravega/internal/database"
	"challenge-fravega/internal/route"
	routePoint "challenge-fravega/internal/route-point"
	"challenge-fravega/internal/vehicle"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// Dependencies
	db := openConnectionDb()

	// Run migrations
	migrationsDir := getEnv("MIGRATIONS_DIR", "./db/migrations")
	if err := database.MigrateDB(db, migrationsDir); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Repositories
	routeRepository := route.NewRepository(db)
	routePointRepository := routePoint.NewRepository(db)
	carDriverRepository := carDriver.NewRepository(db)
	vehicleRepository := vehicle.NewRepository(db)

	// Services
	carDriverService := carDriver.NewService(carDriverRepository)
	vehicleService := vehicle.NewService(vehicleRepository)
	routePointService := routePoint.NewService(routePointRepository)
	routeService := route.NewService(routeRepository)

	// Handlers
	routeHandler := handlers.NewRouteHandler(routeService)
	routePointHandler := handlers.NewRoutePointHandler(routePointService)
	carDriverHandler := handlers.NewCarDriverHandler(carDriverService)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)

	app := gin.Default()

	// Routes
	routeHandler.SetupRoutes(app)
	routePointHandler.SetupRoutes(app)
	carDriverHandler.SetupRoutes(app)
	vehicleHandler.SetupRoutes(app)

	port := getEnv("PORT", "8080")
	if err := app.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func openConnectionDb() *gorm.DB {
	dbPath := getEnv("DB_PATH", "./db/data.sqlite")

	// Ensure directory exists
	dbDir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dbDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create database directory: %v", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Connected to database at %s", dbPath)
	return db
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
