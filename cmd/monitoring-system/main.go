package main

import (
	"log"
	"monitoring-system/config"
	"monitoring-system/internal/api"
	"time"
)

func main() {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ticker := time.NewTicker(time.Duration(config.QueryInterval) * time.Second)

	defer ticker.Stop()

	for range ticker.C {
		log.Println("Querying API...", time.Now())
		response, responseTime, err := api.QueryAPI(config.APIURL, config.APIKey)
		if err != nil {
			log.Printf("Failed to query API: %v", err)
			continue
		}

		apiData := api.ProcessAPIResponse(response, responseTime)
		log.Printf("Processed API Data Response Time: %+v", apiData.ResponseTime)

		if err := api.SendDataToPRTG(config, apiData); err != nil {
			log.Printf("Failed to send data to PRTG: %v", err)
		}
	}
}
