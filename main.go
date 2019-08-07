package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sync"
)

// ResultType is type of measurement result
type ResultType int32

const (
	// AQIResult is an AQI result type
	AQIResult ResultType = iota
	// WeatherResult is an Weather Forecast result type
	WeatherResult
)

// Result is an encapsulation for any type of measurement for a city
type Result struct {
	City        string
	Measurement string
	Type        ResultType
}

// Row is a CSV row representation
type Row struct {
	City        string
	AQI         string
	Temperature string
}

// AsSlice returns the Row as a slice of strings
func (r Row) AsSlice() []string {
	return []string{r.City, r.AQI, r.Temperature}
}

func getAQIMeasurement(cities <-chan []string) chan Result {
	resultChan := make(chan Result)

	go func() {
		for city := range cities {
			aqi, err := getAQI(city[2])
			if err != nil {
				log.Fatalf("Cannot fetch AQI: %s", err)
			}

			result := Result{City: city[2], Type: AQIResult}
			if len(aqi.Results) > 0 {
				result.Measurement = fmt.Sprintf("%f", aqi.Results[0].Measurements[0].Value)
			}
			resultChan <- result
		}
		close(resultChan)
	}()

	return resultChan
}

func getForecast(cities <-chan []string) chan Result {
	resultsChan := make(chan Result)

	go func() {
		for city := range cities {
			resp, err := getDarksky(city[0], city[1])
			if err != nil {
				log.Fatalf("Cannot fetch forecast: %s", err)
			}
			resultsChan <- Result{
				City:        city[2],
				Measurement: fmt.Sprintf("%f", resp.Currently.Temperature),
				Type:        WeatherResult,
			}
		}
		close(resultsChan)
	}()
	return resultsChan
}

func merge(results ...chan Result) chan Result {
	var wg sync.WaitGroup
	wg.Add(len(results))

	outChan := make(chan Result)

	go func() {
		wg.Wait()
		close(outChan)
	}()

	send := func(incoming <-chan Result) {
		for res := range incoming {
			outChan <- res
		}
		wg.Done()
	}

	for _, r := range results {
		go send(r)
	}

	return outChan
}

func main() {
	citiesAQIChan, citiesForecastChan := readCSV("cities.csv")

	aqiChan := getAQIMeasurement(citiesAQIChan)
	forecastsChan := getForecast(citiesForecastChan)

	measurements := make(map[string]Row)

	for result := range merge(aqiChan, forecastsChan) {
		if row, present := measurements[result.City]; present {
			if result.Type == AQIResult {
				row.AQI = result.Measurement
			} else {
				row.Temperature = result.Measurement
			}
		} else {
			row := Row{City: result.City}
			if result.Type == AQIResult {
				row.AQI = result.Measurement
			} else {
				row.Temperature = result.Measurement
			}
			measurements[result.City] = row
		}
	}

	outCSV, err := os.Create("result.csv")
	if err != nil {
		log.Fatalf("Cannot create CSV: %v", err)
	}
	defer outCSV.Close()

	writer := csv.NewWriter(outCSV)
	defer writer.Flush()

	for _, row := range measurements {
		err := writer.Write(row.AsSlice())
		if err != nil {
			log.Fatalf("Cannot write to file: %s", err)
		}
	}

}
