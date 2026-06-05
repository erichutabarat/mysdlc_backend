package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

type turnstileResponse struct {
    Success    bool     `json:"success"`
    ErrorCodes []string `json:"error-codes"` // ← add this
}

func VerifyTurnstile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from header (or change to c.PostForm if sent in body)
		token := c.GetHeader("X-Turnstile-Token")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "missing captcha token"})
			return
		}

		// Verify with Cloudflare
		resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify", url.Values{
			"secret":   {os.Getenv("TURNSTILE_SECRET_KEY")},
			"response": {token},
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "captcha verification failed"})
			return
		}
		defer resp.Body.Close()

		var result turnstileResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil || !result.Success {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "invalid captcha"})
			return
		}
		log.Printf("Turnstile success: %v, errors: %v", result.Success, result.ErrorCodes)
		// Proceed to the next handler
		c.Next()
	}
}