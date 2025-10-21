package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/Michael-Obele/web-scraper-backend/handlers"
	"github.com/Michael-Obele/web-scraper-backend/src/api"
	"github.com/Michael-Obele/web-scraper-backend/src/config"
	"github.com/Michael-Obele/web-scraper-backend/src/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kpechenenko/rword"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Initialize services
	scraperService := services.NewScraperService(cfg)
	defer scraperService.Close() // Ensure Chromedp is closed on exit

	scrapeHandler := api.NewScrapeHandler(scraperService)

	// Initialize rword generator once at startup (fallback to nil on error)
	var wordGen rword.GenerateRandom
	if g, err := rword.New(); err == nil {
		wordGen = g
		log.Println("rword dictionary loaded, using random-word generator")
	} else {
		wordGen = nil
		log.Printf("rword init failed, falling back to simple phrases: %v", err)
	}

	router := gin.Default()

	// CORS middleware configuration
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))

	// Phrase templates that may use 1 or 2 random words
	templates := []string{
		"Your project is live — %s!",
		"All systems go. %s %s",
		"Backend up and running — %s",
		"Happy scraping — %s",
		"Ready when you are: %s %s",
	}

	// Register routes
	router.GET("/", func(c *gin.Context) {
		// If rword is available, produce template-based phrase
		if wordGen != nil {
			// pick a template
			t := templates[rand.Intn(len(templates))]
			// count placeholders (%s)
			count := strings.Count(t, "%s")
			// get that many random words
			words := wordGen.WordList(count)
			// Title-case the first word for readability
			if len(words) > 0 {
				words[0] = strings.Title(words[0])
			}
			// format template with random words
			message := fmt.Sprintf(t, interfaceSlice(words)...)
			log.Printf("GET / - %s", message)
			c.String(http.StatusOK, message)
			return
		}

		// Fallback: use the simple static phrases
		phrases := []string{
			"Your project is live!",
			"All systems go.",
			"Happy scraping!",
		}
		c.String(http.StatusOK, phrases[rand.Intn(len(phrases))])
	})
	router.GET("/health", handlers.HealthCheck)
	router.GET("/scrape", scrapeHandler.HandleScrape)

	// Global handler for unknown routes - log and return a 404 response
	router.NoRoute(func(c *gin.Context) {
		log.Printf("NoRoute: %s %s from %s", c.Request.Method, c.Request.URL.Path, c.ClientIP())
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
	})

	// Custom HTTP server for graceful shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.ServerPort,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Starting server on :%s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the requests it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

// helper to convert []string to []interface{} for fmt.Sprintf
func interfaceSlice(ss []string) []interface{} {
	out := make([]interface{}, len(ss))
	for i, s := range ss {
		out[i] = s
	}
	return out
}
