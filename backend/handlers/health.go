package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck responds with service health status
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":    "healthy",
		"service":   "web-scraper-backend",
		"timestamp": time.Now().Unix(),
	})
}
