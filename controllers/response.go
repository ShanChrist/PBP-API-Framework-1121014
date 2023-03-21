package controllers

import "github.com/gin-gonic/gin"

func sendResponse(c *gin.Context, status int, message string) {
	response := Response{
		Status:  status,
		Message: message,
	}
	c.JSON(status, response)
}

func sendUserResponse(c *gin.Context, status int, message string, userData []User) {
	response := UserResponse{
		Status:  status,
		Message: message,
		Data:    userData,
	}
	c.JSON(status, response)
}
