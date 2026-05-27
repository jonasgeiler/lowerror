package main

import (
	"net/http"
	"os"
	"strconv"
)

type handler struct{}

// ServeHTTP implements the http.Handler interface for our custom handler.
// It responds with a status code and an empty body, based on the request path.
func (*handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// We use `[1:]` to skip the leading "/" when parsing the request path.
	// We also utilize the fact that `http.StatusText` returns an empty string
	// for unknown status codes to make sure it's a known one.
	if statusCode, err := strconv.Atoi(r.URL.Path[1:]); err == nil && http.StatusText(statusCode) != "" {
		// If the path is a valid HTTP status code, respond with that code.
		// For example, "/401" will respond with a "401 Unauthorized" status.
		w.WriteHeader(statusCode)
	} else {
		// Respond to any other path with a "404 Not Found" status.
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	// Just panic if the server fails to start.
	panic(
		// Create the server and start listening and serving.
		(&http.Server{
			// If LOWERROR_ADDRESS is unset, it will default to ":http" (":80").
			Addr: os.Getenv("LOWERROR_ADDRESS"),
			// We use a custom handler instead of a ServeMux.
			Handler: &handler{},
		}).ListenAndServe(),
	)
}
