package config

import (
	"fmt"
	"log"
	"os"
	"mysdlc_backend/internal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	// 1. Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. Build DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// 3. Open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // Optimizes performance
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// config/database.go

	// Pass 1 — create tables without the circular FK
	err = db.AutoMigrate(
		&model.User{},
		&model.SDLC{},
		&model.SDLCSteps{},
		&model.Project{},    // creates projects table (CurrentPhaseID column exists but no FK yet)
		&model.ProjectPhase{}, // creates project_phases table (references projects ✓)
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Pass 2 — now add the FK from projects.current_phase_id → project_phases.id
	// project_phases table now exists so this succeeds
	err = db.Exec(`
		ALTER TABLE projects 
		ADD CONSTRAINT fk_projects_current_phase 
		FOREIGN KEY (current_phase_id) REFERENCES project_phases(id)
		ON DELETE SET NULL
	`).Error
	if err != nil {
		// Don't fatal — constraint may already exist on subsequent runs
		log.Printf("Warning: could not add current_phase FK (may already exist): %v", err)
	}

	log.Println("Database connection established")
	return db
}