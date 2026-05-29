package main

import (
	"mysdlc_backend/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db := config.ConnectDB()

	// Pass 'db' into your Repository/Handler constructors
	// userRepo := repository.NewUserRepository(db)
    
	r := gin.Default()
	r.Run(":8080")
}