# HTTP Server Pattern

Basic HTTP server using Go's `net/http` package, demonstrating routing, handlers, and JSON responses.

## Concepts

- **HTTP Protocol**: Application layer protocol on top of TCP
- **ServeMux**: HTTP request router/multiplexer
- **Handlers**: Functions that handle HTTP requests
- **JSON Encoding**: Marshalling Go structs to JSON
- **HTTP Methods**: GET, POST, PUT, DELETE
- **Status Codes**: 200 OK, 404 Not Found, 500 Error

## Usage

```bash
# Start server
go run main.go

# Test endpoints
curl http://localhost:8080
curl http://localhost:8080/health
curl http://localhost:8080/api/greet?name=Tim
curl -i http://localhost:8080/api/greet?name=Tim  # With headers
```

## Endpoints

| Method | Path | Description | Response |
|--------|------|-------------|----------|
| GET | `/` | Home page | Plain text |
| GET | `/health` | Health check | "OK" |
| GET | `/api/greet?name=X` | JSON greeting | JSON object |

## HTTP Request/Response Flow

```
Client                          Server
  │                               │
  ├─── GET /api/greet?name=Tim ──>│
  │                               │
  │                         ┌─────┴─────┐
  │                         │ Handler   │
  │                         │ - Read    │
  │                         │ - Process │
  │                         │ - Respond │
  │                         └─────┬─────┘
  │                               │
  │<── HTTP/1.1 200 OK ──────────┤
  │    {"message":"Hello, Tim!"} │
  │                               │
```

## Why HTTP?

- **Standardized**: Universal protocol for web services
- **Human-readable**: Easy to debug (unlike binary TCP)
- **Stateless**: Each request independent (scalable)
- **Rich ecosystem**: Libraries, tools, frameworks

## Next Steps

See `03-rest-api-template` for a production-ready REST API with:
- Chi Router with middleware
- Repository Pattern
- Database integration
- Clean architecture
