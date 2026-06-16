package router

import (
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/handler"
	"mysdlc_backend/internal/middleware"
)

func RegisterProjectRoutes(rg *gin.RouterGroup, handler *handler.ProjectHandler) {
	projects := rg.Group("/projects")
	projects.Use(middleware.AuthMiddleware("user"))
	{
		projects.GET("", handler.GetAllProjects)
		projects.GET("/:id", handler.GetProjectByID)
		projects.GET("/:id/members", handler.GetProjectMembers)
		projects.POST("/:id/members", handler.AddProjectMember)
		projects.DELETE("/:id", handler.DeleteProjectByID)
		projects.POST("", handler.CreateProject)
		projects.GET("/:id/phases/:phase_id/tasks", handler.GetTasks)
	}
}