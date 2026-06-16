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
		log.Println("Warning: No .env file found, using system environment variables")
	}

	// 2. Build DSN string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	fmt.Printf("DEBUG: Connecting to DSN: %s\n", dsn)

	// 3. Open connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // Optimizes performance
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Pass 1 — create tables without the circular FK
	err = db.AutoMigrate(
		&model.User{},
		&model.SDLC{},
		&model.SDLCSteps{},
		&model.Project{},
		&model.ProjectPhase{},
		&model.ProjectMember{},
		&model.Task{},
		&model.TaskTemplate{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connection established")
	return db
}