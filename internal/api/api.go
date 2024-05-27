package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"monitoring-system/config"
	"net/http"
	"time"
)

type APIData struct {
	ResponseBody string
	ResponseTime time.Duration
}

func ProcessAPIResponse(response string, responseTime time.Duration) APIData {
	return APIData{
		ResponseBody: response,
		ResponseTime: responseTime,
	}
}

func QueryAPI(apiURL, apiKey string) (string, time.Duration, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	startTime := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(startTime)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", 0, fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), elapsed, nil
}

func SendDataToPRTG(config *config.Config, data APIData) error {
	payload := map[string]interface{}{
		"sensor_name":        config.PRTG.SensorName,
		"sensor_description": config.PRTG.SensorDescription,
		"values": map[string]interface{}{
			"response_time": data.ResponseTime.Seconds(),
		},
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/addsensor", config.PRTG.ServerURL), bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.SetBasicAuth(config.PRTG.Username, config.PRTG.Passhash)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send data to PRTG: %s", resp.Status)
	}

	return nil
}
