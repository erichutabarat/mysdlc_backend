package jwt

import (
    "time"
    "github.com/golang-jwt/jwt/v5"

)

err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
    UserID uint   `json:"user_id"`
    Role   string `json:"role"`
    jwt.RegisteredClaims
}

func GenerateToken(userID uint, role string) (string, error) {
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