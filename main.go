package main

import (
	"encoding/json"	
	"net/http"
)

type Airport struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	IATA    string `json:"iata"`
	ImageURL string `json:"image_url"`
}

type AirportV2 struct {
	Airport
	RunwayLength int `json:"runway_length"`
}

// Mock data for airports in Bangladesh
var airports = []Airport{
	{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"},
	{"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"},
	{"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"},
}

// Mock data for airports in Bangladesh (with runway length for V2)
var airportsV2 = []AirportV2{
	{Airport{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"}, 3200},
	{Airport{"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"}, 2900},
	{Airport{"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"}, 2500},
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: OK"))
}

// Airports handler for the first endpoint
func Airports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airports)
}

// AirportsV2 handler for the second version endpoint
func AirportsV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airportsV2)
}

// ##############################
// ## TODO: Edit this function ##
// ##############################

// UpdateAirportImage handler for updating airport images
func UpdateAirportImage(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the airport name and image data

	// Find the airport by name

	// Initialize GCS client

	// Upload image to GCS and update the airport's image URL

	// Respond with success/failure
}

func main() {
	// Setup routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/airports", Airports)
	http.HandleFunc("/airports_v2", AirportsV2)

	// TODO: complete the UpdateAirportImage handler function
	http.HandleFunc("/update_airport_image", UpdateAirportImage)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
