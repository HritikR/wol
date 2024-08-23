package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config struct to match the expected structure of the JSON config file.
type Config struct {
	MacAddress string `json:"macAddress"`
}

func readConfig() (string, error) {
	// Get the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("failed to get executable path: %w", err)
	}

	// Get the directory containing the executable
	exeDir := filepath.Dir(exePath)

	// Build the full path to the config file in the executable's directory
	configFilePath := filepath.Join(exeDir, "config.json")

	// Read the config file
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", fmt.Errorf("failed to read config file: %w", err)
	}

	// Unmarshal JSON into the Config struct
	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		return "", fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Return the MAC address from the config
	return config.MacAddress, nil
}
