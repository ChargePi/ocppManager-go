# 🔌 ocppManager-go

A library for dynamically managing OCPP configuration (variables). It can read, update, and validate OCPP variables.
Currently, only mandatory key validation is implemented. Value validation will be implemented in the near future.

## ⚡ Usage

Check out the full [example](example/example.go). It also contains a sample configuration file.

``` go
    // Set JSON file format
    manager.SetFileFormat(conf_manager.JSON)

	// Load configuration from file
	err := manager.LoadConfiguration()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Get key value
	value, err := manager.GetConfigurationValue(v16.AuthorizationCacheEnabled.String())
	if err != nil {
		log.Errorf("Error getting configuration value: %v", err)
		return
	}

    // Should be true
	log.Println(value)

	// Update key (only works if the key is can be overwritten)
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
```
