package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	darkskyAPIKey = "API_KEY_HERE"
	darkskyAPIURL = "https://api.darksky.net/forecast/%s/%s,%s"
)

// Currently is the current weather situation
type Currently struct {
	Time                 int     `json:"time"`
	Summary              string  `json:"summary"`
	Icon                 string  `json:"icon"`
	NearestStormDistance int     `json:"nearestStormDistance"`
	PrecipIntensity      float64 `json:"precipIntensity"`
	PrecipProbability    float64 `json:"precipProbability"`
	PrecipType           string  `json:"precipType"`
	PrecipAccumulation   float64 `json:"precipAccumulation"`
	Temperature          float64 `json:"temperature"`
	ApparentTemperature  float64 `json:"apparentTemperature"`
	DewPoint             float64 `json:"dewPoint"`
	Humidity             float64 `json:"humidity"`
	Pressure             float64 `json:"pressure"`
	WindSpeed            float64 `json:"windSpeed"`
	WindGust             float64 `json:"windGust"`
	WindBearing          int     `json:"windBearing"`
	CloudCover           float64 `json:"cloudCover"`
	UvIndex              int     `json:"uvIndex"`
	Visibility           float64 `json:"visibility"`
	Ozone                float64 `json:"ozone"`
}

// DarkskyResponse is the response from Darksky's API
type DarkskyResponse struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timezone  string    `json:"timezone"`
	Currently Currently `json:"currently"`
	Hourly    struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                int     `json:"time"`
			Summary             string  `json:"summary"`
			Icon                string  `json:"icon"`
			PrecipIntensity     float64 `json:"precipIntensity"`
			PrecipProbability   float64 `json:"precipProbability"`
			Temperature         float64 `json:"temperature"`
			ApparentTemperature float64 `json:"apparentTemperature"`
			DewPoint            float64 `json:"dewPoint"`
			Humidity            float64 `json:"humidity"`
			Pressure            float64 `json:"pressure"`
			WindSpeed           float64 `json:"windSpeed"`
			WindGust            float64 `json:"windGust"`
			WindBearing         int     `json:"windBearing"`
			CloudCover          float64 `json:"cloudCover"`
			UvIndex             int     `json:"uvIndex"`
			Visibility          float64 `json:"visibility"`
			Ozone               float64 `json:"ozone"`
			PrecipType          string  `json:"precipType,omitempty"`
			PrecipAccumulation  float64 `json:"precipAccumulation,omitempty"`
		} `json:"data"`
	} `json:"hourly"`
	Daily struct {
		Summary string `json:"summary"`
		Icon    string `json:"icon"`
		Data    []struct {
			Time                        int     `json:"time"`
			Summary                     string  `json:"summary"`
			Icon                        string  `json:"icon"`
			MoonPhase                   float64 `json:"moonPhase"`
			PrecipIntensity             float64 `json:"precipIntensity"`
			PrecipIntensityMax          float64 `json:"precipIntensityMax"`
			PrecipIntensityMaxTime      int     `json:"precipIntensityMaxTime"`
			PrecipProbability           float64 `json:"precipProbability"`
			PrecipType                  string  `json:"precipType"`
			PrecipAccumulation          float64 `json:"precipAccumulation"`
			TemperatureHigh             float64 `json:"temperatureHigh"`
			TemperatureHighTime         int     `json:"temperatureHighTime"`
			TemperatureLow              float64 `json:"temperatureLow"`
			TemperatureLowTime          int     `json:"temperatureLowTime"`
			ApparentTemperatureHigh     float64 `json:"apparentTemperatureHigh"`
			ApparentTemperatureHighTime int     `json:"apparentTemperatureHighTime"`
			ApparentTemperatureLow      float64 `json:"apparentTemperatureLow"`
			ApparentTemperatureLowTime  int     `json:"apparentTemperatureLowTime"`
			DewPoint                    float64 `json:"dewPoint"`
			Humidity                    float64 `json:"humidity"`
			Pressure                    float64 `json:"pressure"`
			WindSpeed                   float64 `json:"windSpeed"`
			WindGust                    float64 `json:"windGust"`
			WindGustTime                int     `json:"windGustTime"`
			WindBearing                 int     `json:"windBearing"`
			CloudCover                  float64 `json:"cloudCover"`
			UvIndex                     int     `json:"uvIndex"`
			UvIndexTime                 int     `json:"uvIndexTime"`
			Visibility                  float64 `json:"visibility"`
			Ozone                       float64 `json:"ozone"`
			TemperatureMin              float64 `json:"temperatureMin"`
			TemperatureMinTime          int     `json:"temperatureMinTime"`
			TemperatureMax              float64 `json:"temperatureMax"`
			TemperatureMaxTime          int     `json:"temperatureMaxTime"`
			ApparentTemperatureMin      float64 `json:"apparentTemperatureMin"`
			ApparentTemperatureMinTime  int     `json:"apparentTemperatureMinTime"`
			ApparentTemperatureMax      float64 `json:"apparentTemperatureMax"`
			ApparentTemperatureMaxTime  int     `json:"apparentTemperatureMaxTime"`
		} `json:"data"`
	} `json:"daily"`
	Flags struct {
		Sources []string `json:"sources"`
		Units   string   `json:"units"`
	} `json:"flags"`
	Offset float64 `json:"offset"`
}

func getDarksky(longitude string, latitude string) (DarkskyResponse, error) {
	var dsResponse DarkskyResponse

	apiURL := fmt.Sprintf(darkskyAPIURL, darkskyAPIKey, latitude, longitude)
	response, err := http.Get(apiURL)

	bodyBytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	err = json.Unmarshal(bodyBytes, &dsResponse)

	return dsResponse, err
}
