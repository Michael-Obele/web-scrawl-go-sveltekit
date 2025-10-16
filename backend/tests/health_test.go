package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Michael-Obele/web-scraper-backend/handlers"
	"github.com/gin-gonic/gin"
)

func TestHealthCheck(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/health", handlers.HealthCheck)

	// Create test request
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	// Execute
	router.ServeHTTP(w, req)

	// Assert status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Assert response structure
	var response map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Errorf("Failed to parse JSON response: %v", err)
	}

	// Assert required fields
	if response["status"] != "healthy" {
		t.Errorf("Expected status='healthy', got %v", response["status"])
	}

	if response["service"] != "web-scraper-backend" {
		t.Errorf("Expected service='web-scraper-backend', got %v", response["service"])
	}

	if _, ok := response["timestamp"].(float64); !ok {
		t.Error("Expected timestamp to be a number")
	}
}
