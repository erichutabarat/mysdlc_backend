package main

import (
	"time"
	"github.com/gin-contrib/cors"
	"mysdlc_backend/internal/config"
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/repository"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/internal/handler"
	"mysdlc_backend/internal/router"
)

// trigger build

func main() {
	// Initialize DB and Gin
	db := config.ConnectDB()
	r := gin.Default()
	r.RedirectTrailingSlash = false
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://127.0.0.1:3000",
			"http://frontend:3000",
			"https://mysdlc.my.id",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Turnstile-Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize repositories, services, handlers
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{Repo: userRepo}
	userHandler := &handler.UserHandler{Service: userService}

	sdlcRepo := &repository.SDLCRepository{DB: db}
	sdlcService := &service.SDLCService{Repo: sdlcRepo}
	sdlcHandler := &handler.SDLCHandler{Service: sdlcService}

	projectRepo := &repository.ProjectRepository{DB: db}
	projectService := &service.ProjectService{Repo: projectRepo, SDLCRepo: sdlcRepo, UserRepo: userRepo}
	projectHandler := &handler.ProjectHandler{Service: projectService}



	// Group API routes
	v1 := r.Group("/api/v1")
	router.RegisterUserRoutes(v1, userHandler)
	router.RegisterSDLCRoutes(v1, sdlcHandler)
	router.RegisterProjectRoutes(v1, projectHandler)
	router.RegisterHealthRoutes(v1, db)

	// Start server
	r.Run("0.0.0.0:8080")
}