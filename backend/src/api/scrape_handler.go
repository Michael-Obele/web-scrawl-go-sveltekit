package api

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/Michael-Obele/web-scraper-backend/src/services"
	"github.com/gin-gonic/gin"
)

// ScrapeHandler handles scrape requests
type ScrapeHandler struct {
	scraperService *services.ScraperService
}

// NewScrapeHandler creates a new scrape handler
func NewScrapeHandler(scraperService *services.ScraperService) *ScrapeHandler {
	return &ScrapeHandler{
		scraperService: scraperService,
	}
}

// HandleScrape handles GET /scrape?url={url}&depth={n}
func (h *ScrapeHandler) HandleScrape(c *gin.Context) {
	// Get URL parameter
	targetURL := c.Query("url")
	if targetURL == "" {
		RespondWithError(c, http.StatusBadRequest, "bad_request", "URL parameter is required")
		return
	}

	// Validate URL
	parsedURL, err := url.ParseRequestURI(targetURL)
	if err != nil || (parsedURL.Scheme != "http" && parsedURL.Scheme != "https") {
		RespondWithError(c, http.StatusBadRequest, "invalid_url", "URL must be a valid HTTP or HTTPS URL")
		return
	}

	// Get depth parameter (default 1)
	depth := 1
	if depthStr := c.Query("depth"); depthStr != "" {
		parsedDepth, err := strconv.Atoi(depthStr)
		if err != nil || parsedDepth < 1 {
			RespondWithError(c, http.StatusBadRequest, "invalid_depth", "Depth must be a positive integer")
			return
		}
		depth = parsedDepth
	}

	// Perform scrape
	result, err := h.scraperService.Scrape(c.Request.Context(), targetURL, depth)
	if err != nil {
		RespondWithError(c, http.StatusInternalServerError, "scrape_failed", err.Error())
		return
	}

	// Return result
	c.JSON(http.StatusOK, result)
}
