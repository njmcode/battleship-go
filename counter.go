package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/davecgh/go-spew/spew" // Linter may complain about this file path :/
)

// CountHandler - has an internal counter var
type CountHandler struct {
	count int

	sync.Mutex // mixin Mutex methods for lock/unlock
}

// CountHandler extends ServeHTTP class
func (c *CountHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	spew.Dump(r)           // debug info to terminal when this handler is called
	count := c.increment() // method defined below

	rw.WriteHeader(200)
	rw.Write([]byte("Count: " + count))
}

// Add an increment method to the CountHandler
func (c *CountHandler) increment() string {
	// Use embedded Mutex functionality so that only one goroutine can
	// call this function at once (to prevent counter going out of sync
	// if multiple clients connect)
	c.Lock()
	defer c.Unlock()

	c.count = c.count + 1
	count := strconv.Itoa(c.count) // convert int to string
	return count
}
