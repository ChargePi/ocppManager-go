package ocpp_v16

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrReadOnly    = errors.New("attribute is read-only")
)

type Key string

func (k Key) String() string {
	return string(k)
}

type Config struct {
	Version int                     `fig:"version" default:"1"`
	Keys    []core.ConfigurationKey `fig:"keys"`
}

// UpdateKey Update the configuration variable in the configuration if it is not readonly.
func (config *Config) UpdateKey(key string, value *string) error {
	log.Debugf("Updating key %s", key)

	// Find the index of the key
	configKey, index, isFound := lo.FindIndexOf(config.Keys, func(item core.ConfigurationKey) bool {
		return item.Key == key
	})
	if !isFound {
		return ErrKeyNotFound
	}

	if configKey.Readonly {
		return ErrReadOnly
	}

	config.Keys[index].Value = value
	return nil
}

// UpdateKeyReadability updates whether the key is updatable or not.
func (config *Config) UpdateKeyReadability(key string, readable bool) error {
	log.Debugf("Updating key readability %s", key)

	// Find the index of the key
	_, index, isFound := lo.FindIndexOf(config.Keys, func(item core.ConfigurationKey) bool {
		return item.Key == key
	})
	if !isFound {
		return ErrKeyNotFound
	}

	config.Keys[index].Readonly = readable
	return nil
}

// GetConfigurationValue Get the value of specified configuration variable in String format.
func (config *Config) GetConfigurationValue(key string) (*string, error) {
	log.Debugf("Getting key %s", key)

	configKey, isFound := lo.Find(config.Keys, func(item core.ConfigurationKey) bool {
		return item.Key == key
	})

	if !isFound {
		return nil, ErrKeyNotFound
	}

	return configKey.Value, nil
}

// GetConfig Get the configuration
func (config *Config) GetConfig() []core.ConfigurationKey {
	return config.Keys
}

// GetVersion Get the current version
func (config *Config) GetVersion() int {
	return config.Version
}

// SetVersion Set the current version
func (config *Config) SetVersion(version int) {
	config.Version = version
}

// Validate validates the configuration - check if all mandatory keys are present.
func (config *Config) Validate(mandatoryKeys []Key) error {
	missingKey := ""

	containsMandatoryKeys := true

	for _, key := range mandatoryKeys {
		containsKey := lo.ContainsBy(config.Keys, func(item core.ConfigurationKey) bool {
			return item.Key == key.String()
		})

		if !containsKey {
			missingKey = strings.Join([]string{missingKey, key.String()}, ", ")
			containsMandatoryKeys = false
		}
	}

	if !containsMandatoryKeys {
		return fmt.Errorf("missing mandatory keys: %s", missingKey)
	}

	return nil
}
