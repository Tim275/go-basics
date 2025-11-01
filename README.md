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
