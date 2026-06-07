package router

import (
	"mysdlc_backend/internal/handler"
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/middleware"
)

func RegisterSDLCRoutes(rg *gin.RouterGroup, h *handler.SDLCHandler){
	sdlc := rg.Group("/sdlc")
	{
		sdlc.GET("", h.GetAllSDLCs)
		sdlc.GET("/:id", h.GetSDLCByID)


		// protected routes for admin only
		authenticated := sdlc.Group("/")
		authenticated.Use(middleware.AuthMiddleware("admin"))
		{
			authenticated.POST("", h.CreateSDLC)
			authenticated.PUT("/:id", h.UpdateSDLC)
			authenticated.DELETE("/:id", h.DeleteSDLC)
			
			// authenticated.POST("/:id/steps", h.CreateSDLCStep)
			// authenticated.PUT("/:id/steps/:id", h.UpdateSDLCStep)
			// authenticated.DELETE("/:id/steps/:id", h.DeleteSDLCStep)
		}
	}
}