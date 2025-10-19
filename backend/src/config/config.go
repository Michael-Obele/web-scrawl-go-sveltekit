package config

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Config holds all configuration values for the scraper service
type Config struct {
	// Server settings
	ServerPort string

	// Scraper settings
	ScraperDelaySeconds    int
	ScraperUserAgents      []string
	ChromedpTimeoutSeconds int
	ScraperTimeoutSeconds  int
	IgnoreRobotsTxt        bool
}

// Load reads configuration from environment variables with sensible defaults
func Load() *Config {
	return &Config{
		ServerPort:          getEnv("PORT", "8080"),
		ScraperDelaySeconds: getEnvAsInt("SCRAPER_DELAY_S", 2),
		ScraperUserAgents: getEnvAsSlice("SCRAPER_USER_AGENTS", []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		}),
		ChromedpTimeoutSeconds: getEnvAsInt("CHROMEDP_TIMEOUT_S", 10),
		ScraperTimeoutSeconds:  getEnvAsInt("SCRAPER_TIMEOUT_S", 30),
		IgnoreRobotsTxt:        getEnvAsBool("SCRAPER_IGNORE_ROBOTS", false),
	}
}

// GetScraperDelay returns the configured delay as a time.Duration
func (c *Config) GetScraperDelay() time.Duration {
	return time.Duration(c.ScraperDelaySeconds) * time.Second
}

// GetChromedpTimeout returns the configured Chromedp timeout as a time.Duration
func (c *Config) GetChromedpTimeout() time.Duration {
	return time.Duration(c.ChromedpTimeoutSeconds) * time.Second
}

// GetScraperTimeout returns the configured scraper timeout as a time.Duration
func (c *Config) GetScraperTimeout() time.Duration {
	return time.Duration(c.ScraperTimeoutSeconds) * time.Second
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsSlice(key string, defaultValue []string) []string {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	return strings.Split(valueStr, ",")
}
