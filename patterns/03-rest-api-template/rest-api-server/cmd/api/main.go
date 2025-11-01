package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq" // PostgreSQL Driver

	"github.com/timour/go-api/internal/db"
	"github.com/timour/go-api/internal/env"
	"github.com/timour/go-api/internal/store"
)

func main() {
	// 1️⃣ Config
	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	// 2️⃣ Database Connection mit db.New()
	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	log.Println("database connection pool established")

	// 3️⃣ Store erstellen
	store := store.NewPostgresStorage(db)

	// 4️⃣ Application erstellen
	app := &application{
		config: cfg,
		store:  store,
	}

	// 5️⃣ Server Setup
	mux := app.mount()

	srv := &http.Server{
		Addr:         cfg.addr,
		Handler:      mux,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Starting server on %s", cfg.addr)

	// 6️⃣ Server starten
	log.Fatal(srv.ListenAndServe())
}
