package services

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/Michael-Obele/web-scraper-backend/src/config"
	"github.com/Michael-Obele/web-scraper-backend/src/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly/v2"
)

// ScraperService handles web scraping operations
type ScraperService struct {
	config      *config.Config
	chromedpCtx context.Context
	cancel      context.CancelFunc
}

// NewScraperService creates a new scraper service and initializes a persistent Chromedp context
func NewScraperService(cfg *config.Config) *ScraperService {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("disable-background-timer-throttling", false),
		chromedp.Flag("disable-backgrounding-occluded-windows", false),
		chromedp.Flag("disable-renderer-backgrounding", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	chromedpCtx, _ := chromedp.NewContext(allocCtx)

	return &ScraperService{
		config:      cfg,
		chromedpCtx: chromedpCtx,
		cancel:      cancel,
	}
}

// Close cleans up the scraper service resources
func (s *ScraperService) Close() {
	s.cancel()
}

// Scrape performs a web scrape of the given URL with the specified depth
func (s *ScraperService) Scrape(ctx context.Context, targetURL string, depth int) (*models.ScrapeResult, error) {
	result := &models.ScrapeResult{
		Links:     []models.Link{},
		Warnings:  []string{},
		FetchedAt: time.Now(),
	}

	// Parse and validate URL
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	// Create a context with timeout
	timeoutCtx, cancel := context.WithTimeout(ctx, s.config.GetScraperTimeout())
	defer cancel()

	// Try to fetch with Chromedp for JS-rendered content first
	html, err := s.fetchWithChromedp(timeoutCtx, targetURL)
	if err != nil {
		log.Printf("Chromedp failed for %s: %v", targetURL, err)
		// Fallback to Colly if Chromedp fails
		result.Warnings = append(result.Warnings, fmt.Sprintf("Chromedp failed (%v), falling back to static fetch", err))
		html, err = s.fetchWithColly(targetURL, depth, result)
		if err != nil {
			log.Printf("Colly also failed for %s: %v", targetURL, err)
			return nil, fmt.Errorf("scraping failed: %w", err)
		}
	} else {
		// Extract links from the Chromedp-rendered HTML
		s.extractLinksFromHTML(html, parsedURL, result)
	}
	if html == "" {
		log.Printf("Warning: Empty HTML captured for %s", targetURL)
		result.Warnings = append(result.Warnings, "Captured HTML was empty")
	}

	// Parse HTML and extract content
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// Extract title
	result.Title = strings.TrimSpace(doc.Find("title").First().Text())
	if result.Title == "" {
		result.Title = parsedURL.Host
	}

	// Convert main content to markdown
	result.Markdown = s.convertToMarkdown(doc)
	result.RawHTML = html

	return result, nil
}

// fetchWithChromedp attempts to fetch content using the persistent headless Chrome instance
func (s *ScraperService) fetchWithChromedp(ctx context.Context, targetURL string) (string, error) {
	// Create a new tab from the persistent browser context
	taskCtx, cancel := chromedp.NewContext(s.chromedpCtx)
	defer cancel()

	// Apply timeout to the tab context
	timeoutCtx, timeoutCancel := context.WithTimeout(taskCtx, s.config.GetChromedpTimeout())
	defer timeoutCancel()

	var html string
	err := chromedp.Run(timeoutCtx,
		chromedp.Navigate(targetURL),
		chromedp.WaitReady("body"),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.OuterHTML("html", &html),
	)

	return html, err
}

// fetchWithColly fetches content using Colly (static crawling)
func (s *ScraperService) fetchWithColly(targetURL string, depth int, result *models.ScrapeResult) (string, error) {
	var html string
	var fetchErr error

	c := colly.NewCollector(
		colly.MaxDepth(depth),
		colly.Async(false),
	)

	// Set user agent rotation
	if len(s.config.ScraperUserAgents) > 0 {
		c.UserAgent = s.config.ScraperUserAgents[0]
	}

	// Apply delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       s.config.GetScraperDelay(),
		RandomDelay: time.Second,
	})

	// Extract links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		text := strings.TrimSpace(e.Text)

		// Resolve relative URLs
		absURL := e.Request.AbsoluteURL(link)
		if absURL != "" {
			result.Links = append(result.Links, models.Link{
				Href: absURL,
				Text: text,
			})
		}
	})

	// Capture HTML
	c.OnResponse(func(r *colly.Response) {
		html = string(r.Body)
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		fetchErr = fmt.Errorf("colly error: %w", err)
	})

	// Visit the URL
	if err := c.Visit(targetURL); err != nil {
		return "", fmt.Errorf("failed to visit URL: %w", err)
	}

	if fetchErr != nil {
		return "", fetchErr
	}

	return html, nil
}

// extractLinksFromHTML extracts links from HTML content
func (s *ScraperService) extractLinksFromHTML(html string, baseURL *url.URL, result *models.ScrapeResult) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return
	}

	doc.Find("a[href]").Each(func(i int, sel *goquery.Selection) {
		href, exists := sel.Attr("href")
		if !exists {
			return
		}

		// Resolve relative URLs
		absURL, err := baseURL.Parse(href)
		if err != nil {
			return
		}

		result.Links = append(result.Links, models.Link{
			Href: absURL.String(),
			Text: strings.TrimSpace(sel.Text()),
		})
	})
}

// convertToMarkdown converts HTML content to markdown
func (s *ScraperService) convertToMarkdown(doc *goquery.Document) string {
	var markdown strings.Builder

	// Process the entire body
	doc.Find("body").Children().Each(func(i int, sel *goquery.Selection) {
		s.nodeToMarkdown(sel, &markdown)
	})

	result := markdown.String()
	if result == "" {
		// Fallback to body text
		result = strings.TrimSpace(doc.Find("body").Text())
	}

	return result
}

func (s *ScraperService) nodeToMarkdown(sel *goquery.Selection, markdown *strings.Builder) {
	// Handle different node types
	switch goquery.NodeName(sel) {
	case "h1", "h2", "h3", "h4", "h5", "h6":
		level := sel.Nodes[0].Data[1] - '0' // Extract heading level from tag name
		text := strings.TrimSpace(sel.Text())
		if text != "" {
			markdown.WriteString(strings.Repeat("#", int(level)) + " " + text + "\n\n")
		}
	case "p":
		text := strings.TrimSpace(sel.Text())
		if text != "" {
			markdown.WriteString(text + "\n\n")
		}
	case "ul", "ol":
		sel.Find("li").Each(func(j int, li *goquery.Selection) {
			text := strings.TrimSpace(li.Text())
			if text != "" {
				if goquery.NodeName(sel) == "ul" {
					markdown.WriteString("- " + text + "\n")
				} else {
					markdown.WriteString(fmt.Sprintf("%d. %s\n", j+1, text))
				}
			}
		})
		markdown.WriteString("\n")
	case "pre":
		code := sel.Find("code")
		lang := code.AttrOr("class", "")
		lang = strings.TrimPrefix(lang, "language-")
		markdown.WriteString(fmt.Sprintf("```%s\n%s\n```\n\n", lang, code.Text()))
	case "code":
		// Handle inline code if not in a pre block
		if sel.Parent().Not("pre").Length() > 0 {
			markdown.WriteString(fmt.Sprintf("`%s`", sel.Text()))
		}
	case "blockquote":
		text := strings.TrimSpace(sel.Text())
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			markdown.WriteString("> " + line + "\n")
		}
		markdown.WriteString("\n")
	case "img":
		alt := sel.AttrOr("alt", "")
		src := sel.AttrOr("src", "")
		markdown.WriteString(fmt.Sprintf("![%s](%s)\n\n", alt, src))
	case "a":
		href := sel.AttrOr("href", "")
		text := sel.Text()
		markdown.WriteString(fmt.Sprintf("[%s](%s)", text, href))
	default:
		// For other block-level elements, just get the text
		if isBlockElement(goquery.NodeName(sel)) {
			text := strings.TrimSpace(sel.Text())
			if text != "" {
				markdown.WriteString(text + "\n\n")
			}
		} else { // Inline elements
			markdown.WriteString(sel.Text())
		}
	}
}

func isBlockElement(tag string) bool {
	blockTags := map[string]bool{
		"address":    true,
		"article":    true,
		"aside":      true,
		"blockquote": true,
		"canvas":     true,
		"dd":         true,
		"div":        true,
		"dl":         true,
		"dt":         true,
		"fieldset":   true,
		"figcaption": true,
		"figure":     true,
		"footer":     true,
		"form":       true,
		"h1":         true,
		"h2":         true,
		"h3":         true,
		"h4":         true,
		"h5":         true,
		"h6":         true,
		"header":     true,
		"hr":         true,
		"li":         true,
		"main":       true,
		"nav":        true,
		"noscript":   true,
		"ol":         true,
		"p":          true,
		"pre":        true,
		"section":    true,
		"table":      true,
		"tfoot":      true,
		"ul":         true,
		"video":      true,
	}
	return blockTags[tag]
}
