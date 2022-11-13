package main

import (
	log "github.com/sirupsen/logrus"
	manager "github.com/xBlaz3kx/ocppManager-go"
	v16 "github.com/xBlaz3kx/ocppManager-go/configuration"
)

func main() {
	log.SetLevel(log.DebugLevel)

	manager.SetFilePath("./configuration.json")

	// Load configuration from file
	err := manager.LoadConfiguration()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Get value
	value, err := manager.GetConfigurationValue(v16.AuthorizationCacheEnabled.String())
	if err != nil {
		log.Errorf("Error getting configuration value: %v", err)
		return
	}

	log.Println(value)

	// Update key
	val := "false"
	err = manager.UpdateKey(v16.AuthorizationCacheEnabled.String(), &val)
	if err != nil {
		log.Errorf("Error updating key: %v", err)
		return
	}

	// Update file
	err = manager.UpdateConfigurationFile()
	if err != nil {
		log.Errorf("Error updating configuration file: %v", err)
	}
}
