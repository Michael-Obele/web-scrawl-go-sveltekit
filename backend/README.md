# Web Scraper Backend

Go backend API server for the web scraper application, built with Gin framework.

## Requirements

- **Go**: 1.21 or higher
- **Dependencies**: Managed via Go modules

## Installation

```bash
# Navigate to backend directory
cd backend

# Download dependencies
go mod download

# Install dependencies (if not already done)
go get -u github.com/gin-gonic/gin
go get -u github.com/gin-contrib/cors
go get -u github.com/gocolly/colly/v2
go get -u github.com/chromedp/chromedp
go get -u github.com/PuerkitoBio/goquery
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
go build -o server main.go

# Run built binary
./server
```

### Testing

```bash
# Run all tests
go test ./... -v

# Run tests with coverage
go test ./... -cover

# Run specific test
go test ./tests -v -run TestHealthCheck
```

## Project Structure

```
backend/
├── main.go              # Entry point and server setup
├── handlers/            # HTTP request handlers
│   └── health.go        # Health check endpoint
├── middleware/          # Gin middleware (CORS, auth, etc.)
├── tests/               # Test files
│   └── health_test.go   # Health endpoint tests
├── go.mod               # Go module dependencies
└── go.sum               # Dependency checksums
```

## API Endpoints

### Health Check

```http
GET /health
```

Response:
```json
{
  "status": "healthy",
  "service": "web-scraper-backend",
  "timestamp": 1697443200
}
```

## Configuration

- **Port**: 8080 (configurable via environment)
- **CORS**: Allows requests from `http://localhost:5173` (frontend)
- **Logging**: Console output with Gin's default logger

## Dependencies

- **gin-gonic/gin**: Web framework
- **gin-contrib/cors**: CORS middleware
- **gocolly/colly/v2**: Web scraping framework
- **chromedp/chromedp**: Headless browser automation
- **PuerkitoBio/goquery**: HTML parsing

## Development

The backend uses:
- Gin framework for routing and middleware
- Go modules for dependency management
- Standard Go testing framework
- RESTful API design patterns

## Troubleshooting

### Port 8080 already in use
```bash
# Find process using port 8080
lsof -ti:8080

# Kill the process
kill $(lsof -ti:8080)
```

### Module import errors
```bash
# Clean module cache
go clean -modcache

# Reinstall dependencies
go mod download
go mod tidy
```

## License

[Add your license here]
