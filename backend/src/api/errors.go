package api

import "github.com/gin-gonic/gin"

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, statusCode int, errorType string, message string) {
	c.JSON(statusCode, ErrorResponse{
		Error:   errorType,
		Message: message,
	})
}
