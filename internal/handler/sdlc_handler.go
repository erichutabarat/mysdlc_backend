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
	if !exists || userRole != "admin" {
		response.Error(c, 403, "Forbidden - Admins only")
		return
	}
	
	var sdlc model.SDLC
	if err := c.ShouldBindJSON(&sdlc); err != nil {
		response.Error(c, 400, "Invalid request body")
		return
	}

	if err := h.Service.CreateSDLC(&sdlc); err != nil {
		response.Error(c, 500, "Failed to create SDLC")
		return
	}
	response.Success(c, 201, "SDLC created successfully", sdlc)
}

func (h *SDLCHandler) GetAllSDLCs(c *gin.Context) {
	sdlcs, err := h.Service.GetAllSDLCs()
	if err != nil {
		response.Error(c, 500, "Failed to retrieve SDLCs")
		return
	}
	response.Success(c, 200, "SDLCs retrieved successfully", sdlcs)
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
	if !exists || userRole != "admin" {
		response.Error(c, 403, "Forbidden - Admins only")
		return
	}
	
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.Error(c, 400, "Invalid SDLC ID")
		return
	}

	var sdlc model.SDLC
	if err := c.ShouldBindJSON(&sdlc); err != nil {
		response.Error(c, 400, "Invalid request body")
		return
	}
	sdlc.ID = uint(id)
	
	if err := h.Service.UpdateSDLC(&sdlc); err != nil {
		response.Error(c, 500, "Failed to update SDLC")
		return
	}
	response.Success(c, 200, "SDLC updated successfully", sdlc)
}

func (h *SDLCHandler) DeleteSDLC(c *gin.Context) {
	userRole, exists := c.Get("role")
	if !exists || userRole != "admin" {
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