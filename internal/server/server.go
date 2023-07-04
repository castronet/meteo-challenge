package server

import (
	// To format text
	"fmt"

	// For string conversions
	"strconv"

	// Http server
	"net/http"

	// HTTP services. Go chi details: https://github.com/go-chi/chi
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	// For rendering web templates
	"github.com/foolin/goview"
)

const (
	defaultPort = 8080
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
	// gvConfig.Funcs = template.FuncMap{"avail": avail}
	goview.Use(goview.New(gvConfig))

	return s
}

// Run the server with the given address and optionally port
func (s *Server) Run(address string, args []string) {
	// Set default port value
	port := defaultPort

	// If given port exists, try to set as server port
	if len(args) > 0 {
		convertedPort, err := strconv.Atoi(args[0])
		if err == nil {
			port = convertedPort
		}
	}

	// Define server address
	s.address = fmt.Sprintf("%s:%d", address, port)

	// Run server
	fmt.Printf("Listening on address %s port %d", address, port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", address, port), s.router)
	if err != nil {
		// Panic error
		fmt.Println("Error while starting server")
		fmt.Println(err)
		return
	}
}

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

	return r
}
