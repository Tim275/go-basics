// HTTP Server Pattern - Application Layer Protocol
// Kapitel 2: Going from TCP to HTTP

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Response struct für JSON Responses
type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func main() {
	// 1️⃣ HTTP ServeMux (Router) erstellen
	mux := http.NewServeMux()

	// 2️⃣ Routes registrieren
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/greet", greetHandler)

	// 3️⃣ HTTP Server mit Timeouts konfigurieren
	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Println("HTTP Server listening on :8080")
	log.Println("Try: http://localhost:8080")
	log.Println("     http://localhost:8080/health")
	log.Println("     http://localhost:8080/api/greet")

	// 4️⃣ Server starten (blockiert hier)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

// homeHandler - Simple Text Response
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Nur GET erlauben
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to Go HTTP Server!\n"))
}

// healthHandler - Health Check Endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

// greetHandler - JSON Response Example
func greetHandler(w http.ResponseWriter, r *http.Request) {
	// Nur GET erlauben
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Query Parameter lesen (?name=John)
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}

	// JSON Response erstellen
	response := Response{
		Message: "Hello, " + name + "!",
		Status:  "success",
	}

	// Content-Type Header setzen
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// JSON encodieren und senden
	json.NewEncoder(w).Encode(response)
}

/*
📚 HTTP Konzepte:

1. HTTP baut auf TCP auf (OSI Model):
   ┌──────────────────┐
   │ Layer 7  ⭐      │ ← Application: HTTP, REST APIs
   │ Application      │   (WIR SIND HIER!)
   ├──────────────────┤
   │ Layer 6          │ ← Presentation: JSON, SSL/TLS
   │ Presentation     │
   ├──────────────────┤
   │ Layer 5          │ ← Session: HTTP Sessions
   │ Session          │
   ├──────────────────┤
   │ Layer 4          │ ← Transport: TCP (zuverlässig)
   │ Transport        │
   ├──────────────────┤
   │ Layer 3          │ ← Network: IP Routing
   │ Network          │
   ├──────────────────┤
   │ Layer 2          │ ← Data Link: Ethernet
   │ Data Link        │
   ├──────────────────┤
   │ Layer 1          │ ← Physical: Network cables
   │ Physical         │
   └──────────────────┘

2. HTTP Request:
   GET /api/greet?name=John HTTP/1.1
   Host: localhost:8080
   User-Agent: curl/7.64.1

3. HTTP Response:
   HTTP/1.1 200 OK
   Content-Type: application/json
   Content-Length: 42

   {"message":"Hello, John!","status":"success"}

4. HTTP Methods:
   - GET: Daten abrufen
   - POST: Daten erstellen
   - PUT: Daten aktualisieren
   - DELETE: Daten löschen
   - PATCH: Teilweise aktualisieren

5. Status Codes:
   - 2xx: Success (200 OK, 201 Created)
   - 3xx: Redirect (301 Moved Permanently)
   - 4xx: Client Error (404 Not Found, 400 Bad Request)
   - 5xx: Server Error (500 Internal Server Error)

🔧 Testen:
   # Server starten:
   go run main.go

   # GET Requests:
   curl http://localhost:8080
   curl http://localhost:8080/health
   curl http://localhost:8080/api/greet
   curl http://localhost:8080/api/greet?name=Tim

   # JSON Response ansehen:
   curl -i http://localhost:8080/api/greet?name=Tim

✅ Vorteile von HTTP über TCP:
   - Standardisierte Struktur (Method, Path, Headers, Body)
   - Status Codes (200, 404, 500)
   - Content-Type Negotiation (JSON, HTML, XML)
   - Built-in Caching Support
   - Human-readable Protocol
*/
