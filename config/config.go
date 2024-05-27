package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	APIURL string `json:"api_url"`
	APIKey string `json:"api_key"`
	PRTG   struct {
		ServerURL         string `json:"server_url"`
		Username          string `json:"username"`
		Passhash          string `json:"passhash"`
		SensorName        string `json:"sensor_name"`
		SensorDescription string `json:"sensor_description"`
	} `json:"prtg"`
	QueryInterval int `json:"query_interval"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println("Failed to read config file")
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		log.Println("Failed to unmarshal config data")
		return nil, err
	}

	return &config, nil
}
