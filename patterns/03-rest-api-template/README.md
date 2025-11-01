# REST API Template Pattern

Production-ready REST API template demonstrating clean architecture, repository pattern, and modern Go best practices.

## Architecture Overview

This template implements a **Clean Layered Architecture**:

```
rest-api-template/
├── cmd/
│   ├── api/              # Entry Point (main.go, server setup)
│   └── migrate/          # Database Migrations
├── internal/
│   ├── store/            # Repository Pattern (Data Access Layer)
│   ├── db/               # Database Connection Pool
│   └── env/              # Environment Variables Helper
├── scripts/              # Database Init Scripts
├── docker-compose.yaml   # PostgreSQL Development Setup
├── .air.toml             # Hot Reload Configuration
└── .envrc                # Environment Variables
```

## Design Patterns

### 1. Repository Pattern
**Purpose**: Separates business logic from data access

```go
// What is possible (Interface)
type Storage struct {
    Posts interface { Create(context.Context, *Post) error }
    Users interface { Create(context.Context, *User) error }
}

// How it's implemented (Concrete)
type PostsStorage struct { db *sql.DB }
func (s *PostsStorage) Create(ctx context.Context, post *Post) error {
    // SQL query here
}
```

**Benefits**:
- ✅ Testable (mock storage in tests)
- ✅ Swappable (PostgreSQL → MongoDB → In-Memory)
- ✅ Clean separation of concerns

### 2. Dependency Injection
**Purpose**: Pass dependencies instead of creating them

```go
type application struct {
    config config
    store  store.Storage  // Injected, not created
}

app := &application{
    config: cfg,
    store:  store.NewPostgresStorage(db),
}
```

### 3. Clean Architecture Layers

```
┌─────────────────────────────────────┐
│  cmd/api (Entry Point)              │ ← main.go, server setup
├─────────────────────────────────────┤
│  Handlers (HTTP Layer)              │ ← api.go, health.go
├─────────────────────────────────────┤
│  Business Logic (Service Layer)     │ ← Future: internal/service/
├─────────────────────────────────────┤
│  Repository (Data Access)           │ ← internal/store/
├─────────────────────────────────────┤
│  Database (Infrastructure)          │ ← internal/db/
└─────────────────────────────────────┘
```

## Key Concepts

### Connection Pool Management
Optimized for production performance:
```go
db.SetMaxOpenConns(30)     // Max concurrent connections
db.SetMaxIdleConns(30)     // Max idle connections
db.SetConnMaxIdleTime(15m) // Close idle after 15min
```

### Chi Router Middleware
```go
r.Use(middleware.Recoverer) // Panic recovery
r.Use(middleware.Logger)    // Request logging
r.Route("/v1", func(r chi.Router) {
    r.Get("/health", app.healthCheckHandler)
})
```

### Environment Variables with Fallbacks
```go
addr := env.GetString("ADDR", ":8080")              // Default :8080
maxConns := env.GetInt("DB_MAX_OPEN_CONNS", 30)    // Default 30
```

## Quick Start

```bash
# 1. Start PostgreSQL
docker compose up -d

# 2. Start API with hot reload
air

# 3. Test health endpoint
curl http://localhost:8080/v1/health
```

## Database Schema (Future)

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT,
    user_id INTEGER REFERENCES users(id),
    tags TEXT[],
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## What's Included

### ✅ Development Tools
- **Air**: Hot reload on file changes
- **Docker Compose**: PostgreSQL container
- **Chi Router**: Fast, idiomatic HTTP router
- **Environment Variables**: Flexible configuration

### ✅ Production-Ready Features
- **Connection Pooling**: Optimized database performance
- **Timeouts**: ReadTimeout, WriteTimeout, IdleTimeout
- **Middleware**: Logging, panic recovery
- **API Versioning**: `/v1` prefix for future compatibility

### ✅ Clean Code Practices
- **Separation of Concerns**: Each layer has one responsibility
- **Testability**: Interfaces enable easy mocking
- **Maintainability**: Clear file structure, one handler per file
- **Scalability**: Stateless design, horizontal scaling ready

## Next Steps

Use this template as starting point for:
- **CRUD Operations**: Users, Posts, Comments
- **Authentication**: JWT tokens, sessions
- **Authorization**: Role-based access control
- **Caching**: Redis integration
- **Testing**: Unit tests, integration tests
- **Deployment**: Docker, Kubernetes, Cloud platforms

## Related Patterns

- `01-tcp-server`: Low-level networking fundamentals
- `02-http-server`: HTTP basics before REST
- `04-database-repository`: Deep dive into repository pattern

## Resources

- [Chi Router Docs](https://github.com/go-chi/chi)
- [PostgreSQL Go Driver](https://github.com/lib/pq)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Repository Pattern](https://martinfowler.com/eaaCatalog/repository.html)
