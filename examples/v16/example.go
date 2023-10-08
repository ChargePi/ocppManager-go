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

	log.Println(value)

	// Update key
	val := "false"
	err = manager.UpdateKey(ocpp_v16.AuthorizationCacheEnabled, &val)
	if err != nil {
		log.Errorf("Error updating key: %v", err)
		return
	}
}
