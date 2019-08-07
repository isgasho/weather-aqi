package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func readCSV(path string) (chan []string, chan []string) {
	aqiChan := make(chan []string)
	forecastChan := make(chan []string)

	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open CSV: %v", err)
	}

	reader := csv.NewReader(csvFile)

	go func(reader *csv.Reader) {
		for {
			var line, error = reader.Read()
			if error == io.EOF {
				csvFile.Close()
				close(aqiChan)
				close(forecastChan)
				break
			} else if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			aqiChan <- line
			forecastChan <- line
		}
	}(reader)

	return aqiChan, forecastChan
}
