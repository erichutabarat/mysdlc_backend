package main

import (
	"mysdlc_backend/internal/config"
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/repository"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/internal/handler"
	"mysdlc_backend/internal/router"
)

func main() {
	// Initialize DB and Gin
	db := config.ConnectDB()
	r := gin.Default()

	// Initialize repositories, services, handlers
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	// Group API routes
	v1 := r.Group("/api/v1")
	router.RegisterUserRoutes(v1, userHandler)

	r.Run(":8080")
}