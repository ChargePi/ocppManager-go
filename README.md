# OCPP variable manager

A library for managing OCPP variables in Go. It provides a simple way to manage OCPP configuration variables, including
getting and setting values, validating values, and enforcing mandatory keys.

## Features

- Configuration versioning
- Mandatory key enforcement
- Custom value validation
- Provides sane default values

## Roadmap

- [x] Configuration versioning
- [x] Custom value validation
- [x] Mandatory key enforcement
- [x] Support for OCPP 1.6
- [ ] Support for OCPP 2.0.1

## Installing

```bash
  go get github.com/ChargePi/ocppManager-go@latest
```

## ⚡ Usage

Check out the full [OCPP 1.6 example](examples/v16/example.go).

```go
package main

import (
	"github.com/ChargePi/ocppManager-go/ocpp_v16"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)

	supportedProfiles := []string{core.ProfileName, smartcharging.ProfileName}
	defaultConfig, err := ocpp_v16.DefaultConfiguration(supportedProfiles...)
	if err != nil {
		log.Errorf("Error getting default configuration: %v", err)
		return
	}

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

## Notes

1. This library is still in development, and the API might change in the future.
2. Storing the configuration is not part of this library. You can use a database, a file, or any other storage mechanism
   to store the configuration by getting the configuration as a map and storing it in your preferred way.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.