package handler

import (
	"mysdlc_backend/internal/model"
	"mysdlc_backend/internal/service"
	"mysdlc_backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserHandler struct { Service *service.UserService }

func (h *UserHandler) Create(c *gin.Context) {
    var req model.RegisterRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 400, "Invalid input: " + err.Error())
        return
    }

    user := model.User{
        Email:    req.Email,
        Name:     req.Name,
        Password: req.Password,
    }

    if err := h.Service.Register(&user); err != nil {
        response.Error(c, 500, "Registration failed")
        return
    }

    user.Password = "" 
    registerResponse := model.RegisterResponse{
        ID: user.ID,
        Email: user.Email,
        Name: user.Name,
        Role: user.Role,
    }
    response.Success(c, 201, "User created", registerResponse)
}

func (h *UserHandler) Login(c *gin.Context) {
    var req model.LoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 400, "Invalid request")
        return
    }

    token, err := h.Service.Login(req.Email, req.Password)
    if err != nil {
        response.Error(c, 401, err.Error())
        return
    }

    response.Success(c, 200, "Login successful", gin.H{"token": token})
}

func (h *UserHandler) Profile(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, 401, "Unauthorized")
        return
    }
    
    user, err := h.Service.Profile(userID.(uint))
    if err != nil {
        response.Error(c, 404, "User not found")
        return
    }
    
    response.Success(c, 200, "User profile", user)
}

func (h *UserHandler) Invitation(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, 401, "Unauthorized")
        return
    }
    invitations, err := h.Service.Invitation(userID.(uint))
    if err != nil {
        response.Error(c, 500, "Failed to fetch invitations")
        return
    }
    response.Success(c, 200, "User invitations", invitations)
}

func (h *UserHandler) RespondInvitation(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        response.Error(c, 401, "Unauthorized")
        return
    }
    
    var req model.RespondInvitationRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 400, "Invalid request")
        return
    }

    if err := h.Service.RespondInvitation(userID.(uint), req); err != nil {
        response.Error(c, 500, "Failed to respond to invitation")
        return
    }

    response.Success(c, 200, "Invitation responded to", nil)
}