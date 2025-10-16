package main

import (
	"log"
	"time"

	"github.com/Michael-Obele/web-scraper-backend/handlers"
	"github.com/Michael-Obele/web-scraper-backend/src/api"
	"github.com/Michael-Obele/web-scraper-backend/src/config"
	"github.com/Michael-Obele/web-scraper-backend/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	scraperService := services.NewScraperService(cfg)
	scrapeHandler := api.NewScrapeHandler(scraperService)

	router := gin.Default()

	// CORS middleware configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(corsConfig))

	// Register routes
	router.GET("/health", handlers.HealthCheck)
	router.GET("/scrape", scrapeHandler.HandleScrape)

	log.Printf("Starting server on :%s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
