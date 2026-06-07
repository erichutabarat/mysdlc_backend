package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mysdlc_backend/internal/handler"
)

func RegisterHealthRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	v1 := rg.Group("/health")
	{
		healthH := &handler.HealthHandler{DB: db}
		v1.GET("", healthH.Check)
	}
}