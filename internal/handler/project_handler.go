package handler

import (
	"strconv"
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ProjectHandler struct {
	Service *service.ProjectService
}

func (h *ProjectHandler) CreateProject(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, 401, "Unauthorized")
        return
    }

    var req model.CreateProjectRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 400, err.Error())
        return
    }

    project, err := h.Service.CreateProject(userID.(uint), &req)
    if err != nil {
        response.Error(c, 500, err.Error())
        return
    }

    response.Success(c, 201, "Project created successfully", project)
}

func (h *ProjectHandler) GetAllProjects(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}

	projects, err := h.Service.GetAllProjects(userID.(uint))
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "Projects retrieved successfully", projects)
}

func (h *ProjectHandler) GetProjectByID(c *gin.Context) {
	paramID := c.Param("id")
	userID, exists := c.Get("userID")
	if !exists {
		response.Error(c, 401, "Unauthorized")
		return
	}
	projectID, err := strconv.Atoi(paramID)
	if err != nil {
		response.Error(c, 400, "Invalid project ID")
		return
	}

	project, phase, err := h.Service.GetProjectByID(uint(projectID), userID.(uint))
	if err != nil {
		response.Error(c, 404, "Project not found")
		return
	}

	response.Success(c, 200, "Project retrieved successfully", map[string]interface{}{
		"project": project,
		"phases":  phase,
	})
}

func (h *ProjectHandler) DeleteProjectByID(c *gin.Context) {
	paramID := c.Param("id")
	projectID, err := strconv.Atoi(paramID)
	if err != nil {
		response.Error(c, 400, "Invalid project ID")
		return
	}
	
	err = h.Service.DeleteProjectByID(uint(projectID))
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}
	
	response.Success(c, 200, "Project deleted successfully", nil)
}