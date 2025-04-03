package database

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gorm.io/gorm"
)

// MigrateDB runs all migrations from the specified directory
func MigrateDB(db *gorm.DB, migrationsDir string) error {
	// Create migration table if it doesn't exist
	err := db.Exec(`CREATE TABLE IF NOT EXISTS migrations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`).Error
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get list of applied migrations
	var appliedMigrations []string
	if err := db.Raw("SELECT name FROM migrations ORDER BY id").Scan(&appliedMigrations).Error; err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	appliedMap := make(map[string]bool)
	for _, m := range appliedMigrations {
		appliedMap[m] = true
	}

	// Get migration files
	files, err := getMigrationFiles(migrationsDir)
	if err != nil {
		return fmt.Errorf("failed to get migration files: %w", err)
	}

	// Run pending migrations
	for _, file := range files {
		migrationName := filepath.Base(file)

		// Skip if already applied
		if appliedMap[migrationName] {
			fmt.Printf("Migration already applied: %s\n", migrationName)
			continue
		}

		// Read migration file
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		// Execute migration within a transaction
		tx := db.Begin()
		if err := tx.Exec(string(content)).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to apply migration %s: %w", migrationName, err)
		}

		// Record migration
		if err := tx.Exec("INSERT INTO migrations (name) VALUES (?)", migrationName).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", migrationName, err)
		}

		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", migrationName, err)
		}

		fmt.Printf("Applied migration: %s\n", migrationName)
	}

	return nil
}

// getMigrationFiles returns a sorted list of SQL migration files
func getMigrationFiles(migrationsDir string) ([]string, error) {
	// Check if directory exists
	if _, err := os.Stat(migrationsDir); os.IsNotExist(err) {
		return nil, fmt.Errorf("migrations directory does not exist: %s", migrationsDir)
	}

	var files []string
	err := filepath.Walk(migrationsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".sql") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Sort files by name to ensure correct order
	sort.Strings(files)

	return files, nil
}
