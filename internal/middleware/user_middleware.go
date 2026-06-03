package middleware

import (
	"net/http"
	"fmt"
	"os"
	"strings"
	"log"
	"mysdlc_backend/internal/model"
	"github.com/gin-gonic/gin"
	jwtlib "github.com/golang-jwt/jwt/v5"
	"mysdlc_backend/pkg/jwt"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodOptions {
            c.Next()
            return
        }
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		log.Printf("DEBUG: Authorization Header received: '%s'", authHeader)

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			return
		}

		tokenString := parts[1]

		claims := &jwt.Claims{}
		
		token, err := jwtlib.ParseWithClaims(tokenString, claims, func(token *jwtlib.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Role check
		if claims.Role != model.UserRole(requiredRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			return
		}

		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		fmt.Printf("DEBUG: Authenticated user ID: %d with role: %s\n", claims.UserID, claims.Role)
		c.Next()
	}
}