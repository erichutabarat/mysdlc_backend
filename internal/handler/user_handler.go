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
    response.Success(c, 201, "User created", user)
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