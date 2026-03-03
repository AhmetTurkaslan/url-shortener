# URL Shortener

A high-performance URL shortening service built with Go, similar to bit.ly. Converts long URLs into short, shareable links with click analytics.

## Features

- Shorten long URLs into unique 8-character codes
- Fast redirects using Redis caching
- Click tracking and access analytics
- RESTful API

## Tech Stack

- **Go** + **Gin** — REST API
- **PostgreSQL** + **GORM** — persistent storage
- **Redis** — caching layer for fast redirects

## Architecture

```
url-shortener/
├── cmd/
│   └── main.go
├── config/
│   ├── database.go
│   └── redis.go
├── internal/
│   ├── handlers/
│   │   └── url_handler.go
│   ├── services/
│   │   └── url_service.go
│   └── models/
│       └── url.go
└── go.mod
```

## How It Works

**Redirect flow:**
1. Request comes in for a short code (e.g. `/abc123`)
2. Redis is checked first — if found, redirect immediately
3. If not in Redis, query PostgreSQL, cache in Redis, then redirect
4. Click counter incremented on every visit

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/shorten` | Create a short URL |
| GET | `/:code` | Redirect to original URL |
| GET | `/:code/stats` | Get link statistics |

### POST /shorten
```json
{
  "long_url": "https://www.example.com/very/long/url"
}
```

### GET /:code/stats
```json
{
  "id": 1,
  "code": "abc123xk",
  "long_url": "https://www.example.com/very/long/url",
  "clicks": 42,
  "created_at": "2026-03-02T19:43:36Z"
}
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL
- Redis

### Installation

```bash
git clone https://github.com/kullaniciadi/url-shortener
cd url-shortener
go mod tidy
```

### Environment Variables

Create a `.env` file in the root directory:

```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=url_shortener
DB_PORT=5432
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
PORT=8080
```

### Run

```bash
go run cmd/main.go
```
