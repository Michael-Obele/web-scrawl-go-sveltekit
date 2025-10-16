package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Michael-Obele/web-scraper-backend/src/api"
	"github.com/Michael-Obele/web-scraper-backend/src/config"
	"github.com/Michael-Obele/web-scraper-backend/src/models"
	"github.com/Michael-Obele/web-scraper-backend/src/services"
	"github.com/gin-gonic/gin"
)

func TestScrapeEndpoint_ExampleCom(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	cfg := config.Load()
	scraperService := services.NewScraperService(cfg)
	scrapeHandler := api.NewScrapeHandler(scraperService)

	router := gin.New()
	router.GET("/scrape", scrapeHandler.HandleScrape)

	// Create test request
	req, err := http.NewRequest("GET", "/scrape?url=https://example.com&depth=1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d. Body: %s", w.Code, w.Body.String())
		return
	}

	// Parse response
	var result models.ScrapeResult
	if err := json.Unmarshal(w.Body.Bytes(), &result); err != nil {
		t.Fatalf("Failed to parse response JSON: %v. Body: %s", err, w.Body.String())
	}

	// Assert required fields
	if result.Title == "" {
		t.Error("Expected non-empty title")
	}

	if result.Markdown == "" {
		t.Error("Expected non-empty markdown content")
	}

	if result.Links == nil {
		t.Error("Expected links array (can be empty but not nil)")
	}

	if result.FetchedAt.IsZero() {
		t.Error("Expected valid fetchedAt timestamp")
	}

	// Check timestamp is recent (within last minute)
	if time.Since(result.FetchedAt) > time.Minute {
		t.Errorf("Expected recent timestamp, got %v", result.FetchedAt)
	}

	// Log result for debugging
	t.Logf("Title: %s", result.Title)
	t.Logf("Markdown length: %d bytes", len(result.Markdown))
	t.Logf("Links count: %d", len(result.Links))
	t.Logf("Warnings: %v", result.Warnings)
}

func TestScrapeEndpoint_InvalidURL(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)
	cfg := config.Load()
	scraperService := services.NewScraperService(cfg)
	scrapeHandler := api.NewScrapeHandler(scraperService)

	router := gin.New()
	router.GET("/scrape", scrapeHandler.HandleScrape)

	tests := []struct {
		name           string
		url            string
		expectedStatus int
	}{
		{"Missing URL", "", http.StatusBadRequest},
		{"Invalid URL", "not-a-url", http.StatusBadRequest},
		{"Non-HTTP URL", "ftp://example.com", http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/scrape?url="+tt.url, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
