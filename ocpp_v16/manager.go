package ocpp_v16

import (
	"fmt"
	"sync"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/samber/lo"
)

type (
	KeyValidator    func(Key Key, value *string) bool
	OnUpdateHandler func(value *string) error

	Manager interface {
		SetMandatoryKeys(mandatoryKeys []Key) error
		GetMandatoryKeys() []Key
		RegisterCustomKeyValidator(KeyValidator)
		ValidateKey(key Key, value *string) error
		UpdateKey(key Key, value *string) error
		OnUpdateKey(key Key, handler OnUpdateHandler) error
		GetConfigurationValue(key Key) (*string, error)
		SetConfiguration(configuration Config) error
		GetConfiguration() ([]core.ConfigurationKey, error)
	}

	ManagerV1 struct {
		supportedProfiles []string
		ocppConfig        *Config
		mandatoryKeys     []Key
		keyValidator      KeyValidator
		onUpdateHandlers  map[Key]OnUpdateHandler
		mu                sync.Mutex
	}
)

func NewV16ConfigurationManager(defaultConfiguration Config, profiles ...string) (*ManagerV1, error) {
	mandatoryKeys := GetMandatoryKeysForProfile(profiles...)

	// Validate default configuration
	err := defaultConfiguration.Validate(mandatoryKeys)
	if err != nil {
		return nil, err
	}

	return &ManagerV1{
		ocppConfig:       &defaultConfiguration,
		mandatoryKeys:    mandatoryKeys,
		onUpdateHandlers: make(map[Key]OnUpdateHandler),
		mu:               sync.Mutex{},
	}, nil
}

// SetConfiguration validates the provided and overwrites the current configuration
func (m *ManagerV1) SetConfiguration(configuration Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Validate the configuration before setting it
	err := configuration.Validate(m.mandatoryKeys)
	if err != nil {
		return err
	}

	m.ocppConfig = &configuration
	return nil
}

// RegisterCustomKeyValidator registers a custom key validator
func (m *ManagerV1) RegisterCustomKeyValidator(validator KeyValidator) {
	m.keyValidator = validator
}

// GetMandatoryKeys returns the mandatory keys for the configuration
func (m *ManagerV1) GetMandatoryKeys() []Key {
	return m.mandatoryKeys
}

// SetMandatoryKeys sets the mandatory keys for the configuration
func (m *ManagerV1) SetMandatoryKeys(mandatoryKeys []Key) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, key := range mandatoryKeys {
		isAlreadyPresent := lo.ContainsBy(m.mandatoryKeys, func(k Key) bool {
			return k.String() == key.String()
		})

		if isAlreadyPresent {
			continue
		}

		m.mandatoryKeys = append(m.mandatoryKeys, key)
	}

	return nil
}

// UpdateKey updates the value of a specific key
func (m *ManagerV1) UpdateKey(key Key, value *string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// Validate the key
	err := m.ValidateKey(key, value)
	if err != nil {
		return err
	}

	// Try to update the key
	err = m.ocppConfig.UpdateKey(key.String(), value)
	if err != nil {
		return err
	}

	// Call the update handler if present
	handler, isFound := m.onUpdateHandlers[key]
	if isFound {
		defer func() {
			err = handler(value)
			if err != nil {
				return
			}
		}()
	}

	return nil
}

// GetConfiguration returns the full current configuration
func (m *ManagerV1) GetConfiguration() ([]core.ConfigurationKey, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfig(), nil
}

// GetConfigurationValue returns the value of a specific key
func (m *ManagerV1) GetConfigurationValue(key Key) (*string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfigurationValue(key.String())
}

// ValidateKey validates a specific key by checking if there is a custom validator registered
func (m *ManagerV1) ValidateKey(key Key, value *string) error {
	if m.keyValidator == nil {
		return nil
	}

	isValid := m.keyValidator(key, value)
	if !isValid {
		return fmt.Errorf("key validation failed for key %s", key)
	}

	return nil
}

// OnUpdateKey registers a function to call after a specific key has been updated.
func (m *ManagerV1) OnUpdateKey(key Key, handler OnUpdateHandler) error {
	_, err := m.ocppConfig.GetConfigurationValue(key.String())
	if err != nil {
		return err
	}

	m.onUpdateHandlers[key] = handler
	return nil
}
