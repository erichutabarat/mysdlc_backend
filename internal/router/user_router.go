package router

import (
	"mysdlc_backend/internal/handler"
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/middleware"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
    users := rg.Group("/users")
    {
        // Public routes
        users.POST("/register", h.Create)
        users.POST("/login", h.Login)

        // Protected routes (sub-group with middleware)
        authenticated := users.Group("/")
        authenticated.Use(middleware.AuthMiddleware("user"))
        {
            authenticated.GET("/profile", h.Profile)
            authenticated.GET("/invitation", h.Invitation)
            authenticated.POST("/invitation", h.RespondInvitation)
        }
    }
}