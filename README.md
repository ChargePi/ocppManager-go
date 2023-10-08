# OCPP variable manager

A library for managing OCPP 1.6 and 2.0.1 variables, with mandatory key and custom value validation.

## Installing

```bash
    go get github.com/xBlaz3kx/ocppManager-go
```

## ⚡ Usage

Check out the full [example](examples/v16/example.go). It also contains a sample configuration file.

```go
package main

import (
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	log "github.com/sirupsen/logrus"
	"github.com/xBlaz3kx/ocppManager-go/ocpp_v16"
)

func main() {
	log.SetLevel(log.DebugLevel)

	supportedProfiles := []string{core.ProfileName, smartcharging.ProfileName}
	defaultConfig := ocpp_v16.DefaultConfiguration(supportedProfiles...)
	manager, err := ocpp_v16.NewV16ConfigurationManager(defaultConfig, supportedProfiles...)

	// Get value
	value, err := manager.GetConfigurationValue(ocpp_v16.AuthorizeRemoteTxRequests)
	if err != nil {
		log.Errorf("Error getting configuration value: %v", err)
		return
	}

	log.Println(*value)

	// Update key
	val := "false"
	err = manager.UpdateKey(ocpp_v16.AuthorizeRemoteTxRequests, &val)
	if err != nil {
		log.Errorf("Error updating key: %v", err)
		return
	}

	// Get value
	value, err = manager.GetConfigurationValue(ocpp_v16.AuthorizeRemoteTxRequests)
	if err != nil {
		log.Errorf("Error getting configuration value: %v", err)
		return
	}

	log.Println(*value)

	// Register custom key validator, which will prevent the key from being updated
	manager.RegisterCustomKeyValidator(func(key ocpp_v16.Key, value *string) bool {
		return key != ocpp_v16.AuthorizeRemoteTxRequests
	})

	// Update key
	val = "true"
	err = manager.UpdateKey(ocpp_v16.AuthorizeRemoteTxRequests, &val)
	if err != nil {
		log.Errorf("Error updating key: %v", err)
		return
	}
}

```
