package server

import (
	"fmt"
	"net/http"

	"github.com/imrany/spindle/internal/scrape"
)

func StartServer(addr string, port int) error {
	// Define the API endpoint
	http.HandleFunc("/scrape", scrape.ScrapeHandler)

	// Determine port from environment variable, default to 8080
	portStr := fmt.Sprintf("%d", port)
	if addr != "" {
		portStr = fmt.Sprintf("%s:%d", addr, port)
	}
	
	fmt.Printf("Server listening on %s\n", portStr)
	return http.ListenAndServe(portStr, nil)
}