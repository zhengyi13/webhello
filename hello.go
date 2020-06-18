// Package main provides a top-level server that does stuff.
package main

import (
	"fmt"
	"net/http"
)

// TODO: Handle multiple URLs, differently.
func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8999", nil)
}

// HelloServer serves a response.
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
