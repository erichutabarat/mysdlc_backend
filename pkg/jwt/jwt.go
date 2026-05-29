package jwt

import (
	"log"
	"os"
	"time"
	"mysdlc_backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var jwtKey []byte

// init() runs automatically when the package is initialized
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET is not set in environment variables")
	}
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   model.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role model.UserRole) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}