package main

import (
	"net/http"
)

func main() {
	// Anything in a file which is in the same dir as this one
	// is considered to already be in scope - no need for explicit imports
	baseHandler := &Handler{}       // instance of class from handler.go
	countHandler := &CountHandler{} // from of class counter.go

	// Define routes
	http.Handle("/", baseHandler)
	http.Handle("/count", countHandler)
	// Can also set plain functions as handlers
	http.HandleFunc("/func", ServeStatus)

	// Start server on localhost 8080
	http.ListenAndServe("0.0.0.0:8080", nil)
}
