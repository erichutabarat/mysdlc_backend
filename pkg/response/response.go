package response

import (
	"github.com/gin-gonic/gin"
	"mysdlc_backend/internal/model"
)

func Success(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, model.Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, err string) {
	c.JSON(code, model.Response{
		Success: false,
		Error:   err,
	})
}