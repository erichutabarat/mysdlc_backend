package jwt

import (
	"os"
	"time"
	"mysdlc_backend/internal/model"
	"github.com/golang-jwt/jwt/v5"
	"errors"
)

var jwtKey []byte

// InitJWT initializes the JWT key from environment variables
func InitJWT() error {
    key := os.Getenv("JWT_SECRET")
    if key == "" {
        return errors.New("JWT_SECRET is not set in environment variables")
    }
    jwtKey = []byte(key)
    return nil
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   model.UserRole `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, role model.UserRole) (string, error) {
	if len(jwtKey) == 0 {
        return "", errors.New("jwtKey not initialized: call InitJWT() first")
    }
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