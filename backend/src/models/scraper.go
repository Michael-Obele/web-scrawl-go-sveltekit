package models

import "time"

// Link represents a hyperlink discovered during scraping
type Link struct {
	Href string `json:"href"`           // Absolute URL
	Text string `json:"text,omitempty"` // Link text or anchor
}

// ScrapeResult represents the output from a scrape job
type ScrapeResult struct {
	Title     string    `json:"title"`              // Page title
	RawHTML   string    `json:"rawHtml"`            // Raw HTML content of the page
	Markdown  string    `json:"markdown"`           // Main content converted to Markdown
	Links     []Link    `json:"links"`              // Discovered links
	Warnings  []string  `json:"warnings,omitempty"` // Optional warnings (robots.txt, fallback, etc.)
	FetchedAt time.Time `json:"fetchedAt"`          // ISO-8601 timestamp
}
