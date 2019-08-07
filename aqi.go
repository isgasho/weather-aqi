package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	aqiURL = "https://api.openaq.org/v1/latest?city=%s&parameter=pm10"
)

// Measurement is measurement
type Measurement struct {
	Parameter       string    `json:"parameter"`
	Value           float64   `json:"value"`
	LastUpdated     time.Time `json:"lastUpdated"`
	Unit            string    `json:"unit"`
	SourceName      string    `json:"sourceName"`
	AveragingPeriod struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	} `json:"averagingPeriod"`
}

// AQIResponse is AQI API response
type AQIResponse struct {
	Meta struct {
		Name    string `json:"name"`
		License string `json:"license"`
		Website string `json:"website"`
		Page    int    `json:"page"`
		Limit   int    `json:"limit"`
		Found   int    `json:"found"`
	} `json:"meta"`
	Results []struct {
		Location     string        `json:"location"`
		City         string        `json:"city"`
		Country      string        `json:"country"`
		Distance     float64       `json:"distance"`
		Measurements []Measurement `json:"measurements"`
		Coordinates  struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"coordinates"`
	} `json:"results"`
}

func getAQI(city string) (AQIResponse, error) {
	var aqi AQIResponse
	response, err := http.Get(fmt.Sprintf(aqiURL, city))

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(bodyBytes, &aqi)

	return aqi, err
}
