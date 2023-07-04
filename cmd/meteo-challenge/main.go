package main

import (
	// For format text
	"fmt"

	// To obtain env vars
	"os"

	// Import HTTP server
	"github.com/castronet/meteo-challenge/internal/server"
)

const (
	defaultListenAddr = "127.0.0.1"
)

func main() {
	// Define server
	fmt.Println("Initializating server")
	s := server.New()

	// Get values for the server
	listenAddr := os.Getenv("LISTEN_ADDRESS")
	if listenAddr == "" {
		listenAddr = defaultListenAddr
	}

	// Run the server
	fmt.Println("Starting server")
	s.Run(defaultListenAddr, []string{})
}
