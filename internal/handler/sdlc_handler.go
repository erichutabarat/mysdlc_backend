package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/pkg/response"
)

type SDLCHandler struct {
	Service *service.SDLCService
}

func (h *SDLCHandler) CreateSDLC(c *gin.Context){
	userRole, exists := c.Get("role")
	
	if !exists || userRole != model.UserRole("admin") {
		response.Error(c, 403, "Forbidden - Admins only")
		return
	}
	
	var sdlc model.SDLC
	if err := c.ShouldBindJSON(&sdlc); err != nil {
		response.Error(c, 400, "Invalid request body")
		return
	}
	userID, existsx := c.Get("userID")
	if !existsx {
		response.Error(c, 400, "User ID not found")
		return
	}

	sdlc.CreatedByID = userID.(uint)

	if err := h.Service.CreateSDLC(&sdlc); err != nil {
		response.Error(c, 500, "Failed to create SDLC")
		return
	}
	response.Success(c, 201, "SDLC created successfully", sdlc)
}

func (h *SDLCHandler) GetAllSDLCs(c *gin.Context) {
    filter := c.Query("filter")
    sdlcs, err := h.Service.GetAllSDLCs()
    if err != nil {
        response.Error(c, 500, "Failed to retrieve SDLCs")
        return
    }

    if filter == "full" {
        response.Success(c, 200, "SDLCs retrieved successfully", sdlcs)
        return
    }

    sdlcResponses := make([]model.SDLCResponse, len(sdlcs))
    for i, sdlc := range sdlcs {
        sdlcResponses[i] = model.ToSDLCResponse(sdlc) 
    }
    response.Success(c, 200, "SDLCs retrieved successfully", sdlcResponses)
}

func (h *SDLCHandler) GetSDLCByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, 400, "Invalid SDLC ID")
		return
	}

	sdlc, err := h.Service.GetSDLCByID(uint(id))
	if err != nil {
		response.Error(c, 404, "SDLC not found")
		return
	}
	response.Success(c, 200, "SDLC retrieved successfully", sdlc)
}

func (h *SDLCHandler) UpdateSDLC(c *gin.Context) {

	userRole, exists := c.Get("role")
	if !exists || userRole != model.UserRole("admin") {
		response.Error(c, 403, "Forbidden - Admins only")
		return
	}

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, 400, "Invalid SDLC ID")
		return
	}

	var req model.UpdateSDLCRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request body")
		return
	}

	updatedSDLC, err := h.Service.UpdateSDLC(uint(id), &req)
	if err != nil {
		response.Error(c, 500, "Failed to update SDLC")
		return
	}

	response.Success(c, 200, "SDLC updated successfully", updatedSDLC)
}

func (h *SDLCHandler) DeleteSDLC(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != model.UserRole("admin") {
		response.Error(c, 403, "Forbidden - Admins only")
		return
	}

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, 400, "Invalid SDLC ID")
		return
	}

	if err := h.Service.DeleteSDLC(uint(id)); err != nil {
		response.Error(c, 500, "Failed to delete SDLC")
		return
	}
	response.Success(c, 200, "SDLC deleted successfully", nil)
}