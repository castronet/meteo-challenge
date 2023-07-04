package main

import (
	// For format text
	"fmt"

	// To obtain env vars
	"os"

	// Import HTTP server
	"github.com/castronet/meteo-challenge/internal/server"

	// Parse command line flags
	"flag"
)

const (
	// Default values for http server
	defaultListenAddr = "127.0.0.1"
	defaultListenPort = "8080"
)

// Main function
func main() {
	// Flags (help in this case)
	var help = flag.Bool("help", false, "Show how to use the challenge")
	flag.Parse()
	// Usage Demo
	if *help {
		fmt.Println("OpenMeteo challenge usage:")
		fmt.Println("\tYou can run the challenge with: $ make run")
		fmt.Println("\tAfter the server starts you should use your browser to access to the web server.")
		fmt.Println("\tIf the http servers starts successfully, it will show you the correct URL no access the challenge.")
		fmt.Println("")
		fmt.Println("\tYou also can build the binary and execute it yourself: $ make build")
		fmt.Println("\tthen you need to start the server with: ./meteo-challenge")
		os.Exit(0)
	}

	// Define server
	fmt.Println("Initializating server")
	s := server.New()

	// Get values for the server
	listenAddr := os.Getenv("LISTEN_ADDRESS")
	if listenAddr == "" {
		listenAddr = defaultListenAddr
	}

	listenPort := os.Getenv("LISTEN_PORT")
	if listenPort == "" {
		listenPort = defaultListenPort
	}

	// Run the server
	fmt.Println("Starting server")
	s.Run(listenAddr, listenPort)
}
