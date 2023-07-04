package server

import (
	// To format text
	"fmt"

	// Http server
	"net/http"

	// HTTP services. Go chi details: https://github.com/go-chi/chi
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	// For rendering web templates
	"github.com/foolin/goview"
)

// Struct to represent the server
type Server struct {
	router  *chi.Mux
	address string
}

// Server constructor
func New() *Server {
	// Define new server
	s := &Server{}

	// Define http server router (with its routes)
	s.router = s.defineRoutes()

	// Configure http templating system
	gvConfig := goview.DefaultConfig
	gvConfig.DisableCache = true // goview disable cache (used for testing purposes)
	goview.Use(goview.New(gvConfig))

	return s
}

// Run the server with the given address and optionally port
func (s *Server) Run(address string, port string) {
	// Save server address on Server struct
	s.address = fmt.Sprintf("%s:%s", address, port)

	// Run server
	fmt.Printf("Listening on address %s port %s\n", address, port)
	fmt.Printf("Open your browser to the following URL: http://%s:%s/\n", address, port)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), s.router)
	if err != nil {
		// If the http server could not start return and exit application
		fmt.Println("Error while starting server")
		fmt.Println(err)
		return
	}
}

// Helper function to define HTTP Routes
// For this test we just used GET method requests, but is possible to define POST, PUT, DELETE
func (s *Server) defineRoutes() *chi.Mux {
	r := chi.NewRouter()

	// Add middleware to web server.
	// Logger used to log requests
	r.Use(middleware.Logger)
	// RealIP used to get real IP address
	r.Use(middleware.RealIP)

	// Define define
	r.Get("/", s.indexHandler)
	r.Get("/temperature/{lat}/{lon}", s.temperatureHandler)

	// Returns the router to be able to chain functions to define different routes
	return r
}
