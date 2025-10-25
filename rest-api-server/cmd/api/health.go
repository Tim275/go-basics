package main

import (
	"net/http"
)

// healthCheckHandler ist ein HTTP-Handler für den Health-Check Endpoint
// Er wird aufgerufen, wenn jemand GET /health aufruft
// w = ResponseWriter (um die Antwort zu schreiben)
// r = Request (enthält alle Informationen über den eingehenden Request)
func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Setze den HTTP-Status-Code auf 200 (OK)
	w.WriteHeader(http.StatusOK)

	// Schreibe "OK" als Response-Body
	// []byte("OK") konvertiert den String in ein Byte-Array
	w.Write([]byte("OK"))
}
