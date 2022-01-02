package main

import (
	log "github.com/sirupsen/logrus"
	manager "github.com/xBlaz3kx/ocppManager-go"
	conf_manager "github.com/xBlaz3kx/ocppManager-go/manager"
	v16 "github.com/xBlaz3kx/ocppManager-go/v16"
)

func main() {
	log.SetLevel(log.DebugLevel)

	// Set JSON file format
	manager.SetFileFormat(conf_manager.JSON)

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
	err = manager.UpdateKey(v16.AuthorizationCacheEnabled.String(), "false")
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
