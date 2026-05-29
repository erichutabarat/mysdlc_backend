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