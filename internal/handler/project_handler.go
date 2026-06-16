package handler

import (
	"strconv"
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/pkg/response"
	"errors"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

	type ProjectHandler struct {
		Service service.ProjectServiceInterface
	}

var ErrUnauthorized = errors.New("unauthorized: only project owner can view members")

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

	_, err = h.Service.DeleteProjectByID(uint(projectID), userID.(uint))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		response.Error(c, 404, "Project not found")
		return
	}

	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "Project deleted successfully", nil)
}

// PROJECT MEMBER HANDLERS
func (h *ProjectHandler) GetProjectMembers(c *gin.Context) {
	projectID, _ := strconv.Atoi(c.Param("id"))
	userID := c.GetUint("userID")

	members, err := h.Service.GetProjectMembers(uint(projectID), userID)
	if err != nil {
		if errors.Is(err, ErrUnauthorized) {
			response.Error(c, 403, "You are not the owner of this project")
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.Error(c, 404, "Project not found")
			return
		}

		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, "Project members fetched", members)
}

func (h *ProjectHandler) AddProjectMember(c *gin.Context) {
	paramID := c.Param("id")
	userID, userexists := c.Get("userID")
	if !userexists {
		response.Error(c, 401, "Unauthorized")
		return
	}
	projectID, err := strconv.Atoi(paramID)
	if err != nil {
		response.Error(c, 400, "Invalid project ID")
		return
	}
	var addMemberReq model.AddProjectMemberRequest

	if err := c.ShouldBindJSON(&addMemberReq); err != nil {
		response.Error(c, 400, "Invalid request body: "+err.Error())
		return
	}

	member, err := h.Service.AddProjectMember(uint(projectID), userID.(uint), &addMemberReq)
	if err != nil {
		if errors.Is(err, ErrUnauthorized) {
			response.Error(c, 403, "You are not the owner of this project")
			return
		}

		response.Error(c, 500, err.Error())
		return
	}
	response.Success(c, 200, "Project member added successfully", member)

}