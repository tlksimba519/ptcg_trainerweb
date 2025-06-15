package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    data,
	})
}

func Error(c *gin.Context, status int, msg string) {
	c.JSON(status, ApiResponse{
		Status:  status,
		Message: msg,
		Data:    nil,
	})
}
