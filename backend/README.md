# Web Scraper Backend

Go backend API server for the web scraper application, built with Gin framework and modern scraping libraries.

## Requirements

- **Go**: 1.21 or higher
- **Dependencies**: Managed via Go modules

## Installation

```bash
# Navigate to backend directory
cd backend

# Download dependencies
go mod download
```

## Run Commands

### Development

```bash
# Run the server
go run main.go

# Server will start on http://localhost:8080
```

### Build

```bash
# Build binary
go build -o scraper main.go

# Run built binary
./scraper
```

### Testing

```bash
# Run all tests
go test ./tests -v

# Run tests with coverage
go test ./tests -cover

# Run smoke test against example.com
go test ./tests -v -run TestScrapeEndpoint_ExampleCom
```

## Project Structure

```
backend/
├── main.go                    # Entry point and server setup
├── handlers/                  # HTTP request handlers
│   └── health.go             # Health check endpoint
├── src/                      # Main application code
│   ├── api/                  # API handlers and middleware
│   │   ├── errors.go         # Error handling utilities
│   │   └── scrape_handler.go # Scrape endpoint handler
│   ├── config/               # Configuration management
│   │   └── config.go         # Environment-based config
│   ├── models/               # Data models
│   │   └── scraper.go        # ScrapeResult and Link types
│   └── services/             # Business logic
│       └── scraper_service.go # Core scraping service
├── tests/                    # Test files
│   └── smoke_test.go         # Integration tests
├── go.mod                    # Go module dependencies
└── go.sum                    # Dependency checksums
```

## API Endpoints

### Health Check

```http
GET /health
```

Response:
```json
{
  "status": "ok",
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### Web Scraping

```http
GET /scrape?url={url}&depth={depth}
```

**Parameters:**
- `url` (required): The URL to scrape (must be valid HTTP/HTTPS)
- `depth` (optional): Crawl depth (default: 1, max: 3)

**Success Response (200):**
```json
{
  "title": "Example Domain",
  "rawHtml": "<!doctype html>\n<html>\n<head>\n...",
  "markdown": "# Example Domain\n\nThis domain is for use...",
  "links": [
    {
      "href": "https://www.iana.org/domains/example",
      "text": "More information..."
    }
  ],
  "fetchedAt": "2024-01-01T00:00:00Z",
  "warnings": []
}
```

**Error Responses:**
- `400 Bad Request`: Invalid URL or parameters
- `403 Forbidden`: Access denied (future robots.txt implementation)
- `500 Internal Server Error`: Scraping failed

## Configuration

Environment variables (with defaults):

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | `8080` | Server port |
| `SCRAPER_DELAY_S` | `2` | Delay between requests (seconds) |
| `SCRAPER_USER_AGENTS` | Multiple defaults | Comma-separated user agents |
| `CHROMEDP_TIMEOUT_S` | `10` | Headless browser timeout (seconds) |
| `SCRAPER_TIMEOUT_S` | `30` | Overall scrape timeout (seconds) |
| `SCRAPER_IGNORE_ROBOTS` | `false` | Ignore robots.txt (future feature) |

Example:
```bash
export SCRAPER_DELAY_S=5
export SCRAPER_TIMEOUT_S=60
go run main.go
```

## Scraping Architecture

The backend uses a multi-layered approach for robust web scraping:

### 1. Primary: Headless Chrome (Chromedp)
- Renders JavaScript-heavy pages
- Captures dynamic content
- Handles SPAs and client-side rendering

### 2. Fallback: Static Crawling (Colly)
- Traditional HTTP requests
- Fast for static content
- Includes warning when fallback is used

### 3. Content Processing (GoQuery)
- HTML parsing and cleaning
- Markdown conversion
- Link extraction and normalization

## Dependencies

- **gin-gonic/gin**: Web framework and routing
- **gin-contrib/cors**: CORS middleware for frontend integration
- **gocolly/colly/v2**: Web crawling framework
- **chromedp/chromedp**: Headless browser automation
- **PuerkitoBio/goquery**: HTML parsing and manipulation
- **sirupsen/logrus**: Structured logging (future enhancement)

## Development

### Code Organization

- **Handlers**: HTTP request/response logic
- **Services**: Business logic and external integrations
- **Models**: Data structures and validation
- **Config**: Environment-based configuration
- **Tests**: Unit and integration tests

### Error Handling

- Structured error responses
- Comprehensive logging
- Graceful degradation (Chromedp → Colly fallback)
- Context cancellation for timeouts

### Security Considerations

- Input validation and sanitization
- CORS configuration for frontend origin
- Rate limiting preparation (delay configuration)
- Robots.txt compliance preparation

## Testing Strategy

### Unit Tests
- Individual function testing
- Mock external dependencies
- Fast execution

### Integration Tests
- Full API endpoint testing
- Real HTTP requests to example.com
- End-to-end validation

### Smoke Tests
```bash
# Test basic functionality
go test ./tests -v -run TestScrapeEndpoint_ExampleCom
```

## Troubleshooting

### Port 8080 already in use
```bash
# Find process using port 8080
lsof -ti:8080

# Kill the process
kill $(lsof -ti:8080)
```

### Chromedp issues (headless browser)
```bash
# Check if Chrome/Chromium is installed
which google-chrome || which chromium-browser

# Install on Ubuntu/Debian
sudo apt-get install chromium-browser
```

### Module import errors
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod download
go mod tidy
```

### High memory usage
```bash
# Chromedp can be memory-intensive
# Reduce concurrency or add resource limits
export CHROMEDP_TIMEOUT_S=5
```

## Performance

- **Typical scrape time**: 2-10 seconds for static sites
- **Memory usage**: ~50-200MB per concurrent scrape
- **Concurrent requests**: Limited by system resources
- **Rate limiting**: Configurable delays between requests

## Future Enhancements

- **Robots.txt compliance**: Automatic checking and respect
- **User-agent rotation**: Cycle through configured agents
- **Multi-depth crawling**: Follow links beyond depth=1
- **Content filtering**: Remove ads, navigation, footers
- **Caching**: Store and reuse scraped content
- **Metrics**: Prometheus/Grafana integration

## License

[Add your license here]