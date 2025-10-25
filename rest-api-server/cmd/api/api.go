package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// application struct hält alle Abhängigkeiten für unsere API
// Hier können wir später weitere Dinge hinzufügen wie Datenbanken, Logger, etc.
type application struct {
	config config
}

// config struct enthält alle Konfigurationseinstellungen für den Server
type config struct {
	addr string // Server-Adresse und Port (z.B. ":8000")
}

// mount() registriert alle HTTP-Routen (Endpoints) für unsere API
// Chi Router ist OPTIONAL - Standard http.ServeMux funktioniert auch!
// Aber Chi bietet: Middleware, besseres Routing, API-Versioning
func (app *application) mount() http.Handler {
	// Erstelle einen neuen Chi Router
	r := chi.NewRouter()

	// Middleware - wird für ALLE Requests ausgeführt
	r.Use(middleware.Recoverer) // Panic Recovery
	r.Use(middleware.Logger)    // Request Logging

	// API Version 1 Route-Gruppe
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
		// Später weitere Endpoints:
		// r.Get("/users", app.getUsersHandler)
		// r.Post("/users", app.createUserHandler)
	})

	return r
}

// run() startet den HTTP-Server und konfiguriert ihn
// Es gibt einen error zurück, falls der Server nicht starten kann
func (app *application) run() error {
	// Hole alle registrierten Routen vom mount() Methode
	mux := app.mount()

	// Erstelle einen neuen HTTP-Server mit erweiterten Einstellungen
	srv := &http.Server{
		Addr:         app.config.addr,  // Server-Adresse (z.B. ":8000")
		Handler:      mux,              // Router mit allen Routen
		WriteTimeout: 30 * time.Second, // Max. Zeit zum Senden einer Response
		ReadTimeout:  10 * time.Second, // Max. Zeit zum Lesen eines Requests
		IdleTimeout:  time.Minute,      // Max. Zeit für Keep-Alive Verbindungen
	}

	// Logge eine Nachricht, dass der Server startet
	log.Printf("Starting server on %s", app.config.addr)

	// Starte den Server und höre auf eingehende HTTP-Requests
	// ListenAndServe() blockiert hier und läuft bis der Server beendet wird
	return srv.ListenAndServe()
}
