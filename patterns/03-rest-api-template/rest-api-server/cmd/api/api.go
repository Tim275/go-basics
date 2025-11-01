package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/timour/go-api/internal/store"
)

// application struct hält alle Abhängigkeiten für unsere API
type application struct {
	config config
	store  store.Storage
}

// config struct enthält alle Konfigurationseinstellungen
type config struct {
	addr string   // Server-Adresse und Port
	db   dbConfig // Database Configuration
}

// dbConfig enthält Database Connection Pool Settings
type dbConfig struct {
	addr         string // Database Connection String
	maxOpenConns int    // Max. offene Connections
	maxIdleConns int    // Max. idle Connections
	maxIdleTime  string // Max. idle Time (z.B. "15m")
}

// mount() registriert alle HTTP-Routen (Endpoints) für unsere API
func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer) // Panic Recovery
	r.Use(middleware.Logger)    // Request Logging

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}
