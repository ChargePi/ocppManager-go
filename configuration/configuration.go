package configuration

import (
	"errors"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	log "github.com/sirupsen/logrus"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrReadOnly    = errors.New("attribute is read-only")
)

type (
	Key string

	Config struct {
		Version int                     `fig:"version" default:"1"`
		Keys    []core.ConfigurationKey `fig:"keys"`
	}
)

func (k Key) String() string {
	return string(k)
}

// UpdateKey Update the configuration variable in the configuration if it is not readonly.
func (config *Config) UpdateKey(key string, value *string) error {
	log.Debugf("Updating key %s", key)

	for i, configKey := range config.Keys {
		if configKey.Key == key {
			if !configKey.Readonly {
				config.Keys[i].Value = value
				return nil
			}

			return ErrReadOnly
		}
	}

	return ErrKeyNotFound
}

//GetConfigurationValue Get the value of specified configuration variable in String format.
func (config *Config) GetConfigurationValue(key string) (*string, error) {
	log.Debugf("Getting key %s", key)

	for _, configKey := range config.Keys {
		if configKey.Key == key {
			return configKey.Value, nil
		}
	}

	return nil, ErrKeyNotFound
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
