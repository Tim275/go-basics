package main

import (
	"log"
	"os"
)

func main() {
	// Konfiguration aus Environment Variables laden
	cfg := config{
		addr: getEnv("ADDR", ":8000"), // Default: :8000
	}

	// Application-Instanz erstellen
	app := &application{
		config: cfg,
	}

	// Server starten
	log.Fatal(app.run())
}

// getEnv holt Environment Variable oder gibt Default zurück
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
