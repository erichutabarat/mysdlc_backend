package router

import (
	"mysdlc_backend/internal/handler"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, h *handler.UserHandler) {
	users := rg.Group("/users")
	{
		users.POST("/register", h.Create)
		users.POST("/login", h.Login)
	}
}