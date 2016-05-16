package lib

import (
	"encoding/json"
	"os"
)

// This the one and only Configuration-object.
var config *Configuration

// The whole configuration.
type Configuration struct {
	ApplicationUrl string
}

// Reads a configuration-file from a specified path.
func ReadConfigurationFromFile(filePath string) {
	file, _ := os.Open(filePath)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&config)
	if err != nil {
		println("Unable to read configuration-file.")
		os.Exit(1)
	}
}

// Returns the current Configuration-object.
func GetConfiguration() *Configuration {
	if config == nil {
		println("No active configuration found.")
		os.Exit(1)
	} else {
		return config
	}
	return nil
}
