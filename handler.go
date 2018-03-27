package main

import (
	"net/http"
)

// Handler - no internal props of its own
type Handler struct{}

// Handler extends the ServeHTTP class
func (*Handler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)    // Set response status code, must be done first
	rw.Write([]byte("OK")) // Return byte stream form of 'OK' string
}

// ServeStatus - regular function (i.e. not an extended class)
// which can also act as a route handler (see main.go)
func ServeStatus(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("Also OK"))
}
