package config_test

import (
	"os"
	"testing"
	"github.com/joho/godotenv"
)

func TestDB(t *testing.T) {

	t.Run("Read Config DB", func(t *testing.T) {
		err := godotenv.Load("../../.env.dev")
		if err != nil {
			t.Fatalf("Failed to load .env.test file: %v", err)
		}
		dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
		
		if dsn == "" {
			t.Fatal("DSN is empty, check if environment variables are set correctly")
		}

		expectedDSN := "user:pass@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
		if dsn != expectedDSN {
			t.Errorf("Expected DSN '%s', got '%s'", expectedDSN, dsn)
		}
	})
}