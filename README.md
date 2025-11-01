# Go Basics - Learning Go Backend Engineering

A comprehensive collection of Go patterns, concepts, and templates for building production-ready backend systems.

## 📚 What's Inside

This repository is organized into **patterns** (code templates) and **concepts** (advanced topics):

```
go-basics/
├── patterns/               # Reusable code templates
│   ├── 01-tcp-server/     # TCP networking fundamentals
│   ├── 02-http-server/    # HTTP basics with net/http
│   ├── 03-rest-api-template/  # Production REST API template
│   └── 04-database-repository/ # Repository pattern deep dive
├── advanced-go-concepts/  # Advanced Go programming
│   ├── Goroutines/
│   ├── Channels/
│   ├── Context_and_Timeouts/
│   ├── Testing/
│   └── ...
└── docs/                  # Concept explanations
```

## 🚀 Quick Start

### Prerequisites
- Go 1.25+
- Docker & Docker Compose (for database patterns)
- Air (optional, for hot reload): `go install github.com/air-verse/air@latest`

### Try a Pattern

```bash
# 1. Clone the repository
git clone https://github.com/Tim275/go-basics.git
cd go-basics

# 2. Run TCP Server
cd patterns/01-tcp-server
go run main.go

# 3. Run HTTP Server
cd ../02-http-server
go run main.go

# 4. Run REST API Template
cd ../03-rest-api-template
docker compose up -d  # Start PostgreSQL
air                   # Start with hot reload
curl http://localhost:8080/v1/health
```

## 📖 Learning Path

### Level 1: Networking Fundamentals
Start here if you're new to network programming:

1. **[01-tcp-server](patterns/01-tcp-server/)** - TCP basics, goroutines, concurrent connections
   - Learn: `net.Listen()`, `net.Accept()`, goroutines per connection
   - Practice: Build an echo server, chat server

2. **[02-http-server](patterns/02-http-server/)** - HTTP protocol, routing, JSON
   - Learn: `http.ServeMux`, handlers, JSON encoding/decoding
   - Practice: Build API endpoints, query parameters

### Level 2: Production Patterns
Ready for real-world applications:

3. **[03-rest-api-template](patterns/03-rest-api-template/)** - Complete REST API
   - Learn: Clean architecture, Chi router, middleware, connection pooling
   - Practice: Build a full CRUD API with PostgreSQL

4. **[04-database-repository](patterns/04-database-repository/)** - Repository Pattern
   - Learn: Data access abstraction, testing strategies, SQL best practices
   - Practice: Implement CRUD operations, transactions

### Level 3: Advanced Concepts
Deep dive into Go's powerful features:

5. **[Advanced Go Concepts](advanced-go-concepts/)** - Concurrency, testing, optimization
   - Goroutines & Channels
   - Context & Timeouts
   - Worker Pool Pattern
   - Mutexes & Race Conditions
   - Testing & Benchmarking
   - Error Handling Patterns

## 🎯 Key Concepts

### Network Layers (OSI Model)
```
┌────────────────────┐
│ Layer 7            │ ← Application: HTTP, FTP, SMTP, DNS
│ Application        │   (User services, APIs, Web)
├────────────────────┤
│ Layer 6            │ ← Presentation: Encryption, Compression
│ Presentation       │   (SSL/TLS, JPEG, ASCII, JSON)
├────────────────────┤
│ Layer 5            │ ← Session: Connection Management
│ Session            │   (Login sessions, API sessions)
├────────────────────┤
│ Layer 4            │ ← Transport: TCP (reliable), UDP (fast)
│ Transport          │   (Ports, Segments, Flow control)
├────────────────────┤
│ Layer 3            │ ← Network: IP Routing, Addressing
│ Network            │   (IP addresses, Routers, Packets)
├────────────────────┤
│ Layer 2            │ ← Data Link: MAC addresses, Switching
│ Data Link          │   (Ethernet, WiFi, Frames)
├────────────────────┤
│ Layer 1            │ ← Physical: Cables, Signals, Bits
│ Physical           │   (Fiber optic, Copper, Radio waves)
└────────────────────┘
```

**Our Focus**: Layers 4 (TCP) and 7 (HTTP/REST APIs)

### Clean Architecture Layers
```
Entry Point (cmd/)
    ↓
HTTP Handlers (API Layer)
    ↓
Business Logic (Service Layer)
    ↓
Repository (Data Access)
    ↓
Database (Infrastructure)
```

### Go Concurrency Model
```
Main Goroutine
    ├─ Worker 1 (goroutine)
    ├─ Worker 2 (goroutine)
    └─ Worker 3 (goroutine)
         ↓
    Channels (communication)
```

## 🛠️ Patterns Covered

### Design Patterns
- **Repository Pattern**: Clean data access abstraction
- **Dependency Injection**: Testable, maintainable code
- **Middleware Pattern**: Cross-cutting concerns (logging, auth)
- **Worker Pool**: Controlled concurrent processing

### Architectural Patterns
- **Clean Architecture**: Layered separation of concerns
- **REST API Design**: Resource-based HTTP endpoints
- **Connection Pooling**: Optimized database performance

### Go-Specific Patterns
- **Goroutines for Concurrency**: Lightweight thread model
- **Channels for Communication**: Safe data sharing
- **Context for Cancellation**: Timeout and deadline handling
- **Interfaces for Abstraction**: Duck typing, testability

## 📝 Code Examples

### TCP Server (5 lines)
```go
listener, _ := net.Listen("tcp", ":8080")
for {
    conn, _ := listener.Accept()
    go handleConnection(conn)  // Concurrent!
}
```

### HTTP Server (3 lines)
```go
http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
```

### Repository Pattern (Clean!)
```go
// Business logic is clean
err := app.store.Users.Create(ctx, &user)
```

## 🎓 Real-World Application

All patterns are used in production applications like:
- **[GopherSocial API](https://github.com/Tim275/social-media-api)** - Full social media backend
  - 21-chapter curriculum implementation
  - Posts, Users, Followers, Auth, Caching
  - PostgreSQL, Redis, Swagger docs

## 📚 Resources

### Official Go Documentation
- [Go Tour](https://go.dev/tour/) - Interactive introduction
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
- [Go by Example](https://gobyexample.com/) - Code examples

### Network Programming
- [TCP RFC 793](https://tools.ietf.org/html/rfc793) - TCP specification
- [HTTP/1.1 RFC 2616](https://tools.ietf.org/html/rfc2616) - HTTP spec

### Architecture & Patterns
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) - Uncle Bob
- [Repository Pattern](https://martinfowler.com/eaaCatalog/repository.html) - Martin Fowler

### Go Libraries Used
- [Chi Router](https://github.com/go-chi/chi) - Lightweight HTTP router
- [pq](https://github.com/lib/pq) - PostgreSQL driver
- [Air](https://github.com/air-verse/air) - Hot reload

## 🤝 Contributing

This is a learning repository. Feel free to:
- Report issues or unclear explanations
- Suggest improvements to patterns
- Add new patterns or examples

## 📄 License

MIT License - Free to use for learning and commercial projects

## 🚀 Next Steps

1. **Start with patterns**: Work through 01 → 04 sequentially
2. **Build a project**: Use `03-rest-api-template` as starter
3. **Learn advanced concepts**: Explore `advanced-go-concepts/`
4. **Real application**: Check out [GopherSocial API](https://github.com/Tim275/social-media-api)

---

**Happy Learning!** 🎉

For questions or feedback, open an issue or reach out
