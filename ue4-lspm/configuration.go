package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ProcessManagerServerConfiguration is used to hold configuration supplied to the server.
type ProcessManagerServerConfiguration struct {
	DedicatedServerPath string `json:"dedicated_server_path"`
	TotalProcesses      int    `json:"total_processes"`
}

// NewProcessManagerServerConfiguration creates a new Configuration object and returns the pointer to the configuration.
// This configuration is marshalled directly from a supplied JSON file.
func NewProcessManagerServerConfiguration(path string) (*ProcessManagerServerConfiguration, error) {
	configurationObject := &ProcessManagerServerConfiguration{}
	configuration, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error reading")
		return nil, err
	}
	err = json.Unmarshal(configuration, configurationObject)
	if err != nil {
		return nil, err
	}
	return configurationObject, nil
}
