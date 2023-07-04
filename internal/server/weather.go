package server

import (
	// Used to format text
	"fmt"

	// Used to fetch data from API
	"net/http"

	// to manipulate data
	"io/ioutil"

	// JSON encoding
	"encoding/json"

	// For rendering web templates
	"github.com/foolin/goview"

	// HTTP Services
	"github.com/go-chi/chi/v5"
)

// Struct to define API response
type OpenMeteo struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time          string `json:"time"`
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

// Function to handle GET requests on "/" endpoint
func (s *Server) indexHandler(w http.ResponseWriter, r *http.Request) {
	err := goview.Render(w, http.StatusOK, "index", goview.M{
		"WebTitle":  "Index",
		"serverURL": s.address,
	})

	if err != nil {
		fmt.Printf("ERROR: Render index.html error: %v!\n", err)
	}
}

// Function to handle GET requests on "/temperature/lat/lon" endpoint
func (s *Server) temperatureHandler(w http.ResponseWriter, r *http.Request) {
	lat := chi.URLParam(r, "lat")
	lon := chi.URLParam(r, "lon")

	apiUrl := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&hourly=temperature_2m", lat, lon)
	openMeteo := getTemperature(apiUrl)

	err := goview.Render(w, http.StatusOK, "temperature", goview.M{
		"WebTitle": fmt.Sprintf("Temperature on %.2f,%.2f", openMeteo.Latitude, openMeteo.Longitude),
		"data":     openMeteo,
	})

	if err != nil {
		fmt.Printf("ERROR: Render temperature.html error: %v!\n", err)
	}
}

// Helper function to catch data from OpenMeteo API
func getTemperature(apiUrl string) OpenMeteo {
	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Printf("ERROR: Unable to fetch data from OpenMeteo: %v!\n", err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("ERROR: Unable to read data from body from OpenMeteo: %v!\n", err)
	}

	var OpenMeteoObject OpenMeteo
	json.Unmarshal(responseData, &OpenMeteoObject)

	return OpenMeteoObject
}
