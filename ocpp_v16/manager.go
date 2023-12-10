package ocpp_v16

import (
	"fmt"
	"sync"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/samber/lo"
)

type (
	KeyValidator func(Key Key, value *string) bool

	Manager interface {
		SetMandatoryKeys(mandatoryKeys []Key) error
		GetMandatoryKeys() []Key
		RegisterCustomKeyValidator(KeyValidator)
		UpdateKey(key Key, value *string) error
		GetConfigurationValue(key Key) (*string, error)
		SetConfiguration(configuration Config) error
		GetConfiguration() ([]core.ConfigurationKey, error)
	}

	ManagerImpl struct {
		supportedProfiles []string
		ocppConfig        *Config
		mandatoryKeys     []Key
		keyValidator      KeyValidator
		mu                sync.Mutex
	}
)

func NewV16ConfigurationManager(defaultConfiguration Config, profiles ...string) (*ManagerImpl, error) {
	mandatoryKeys := GetMandatoryKeysForProfile(profiles...)

	// Validate default configuration
	err := validateConfiguration(defaultConfiguration, mandatoryKeys)
	if err != nil {
		return nil, err
	}

	return &ManagerImpl{
		ocppConfig:    &defaultConfiguration,
		mandatoryKeys: mandatoryKeys,
		mu:            sync.Mutex{},
	}, nil
}

func (m *ManagerImpl) SetConfiguration(configuration Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	err := validateConfiguration(configuration, m.mandatoryKeys)
	if err != nil {
		return err
	}

	m.ocppConfig = &configuration
	return nil
}

func (m *ManagerImpl) RegisterCustomKeyValidator(validator KeyValidator) {
	m.keyValidator = validator
}

func (m *ManagerImpl) GetMandatoryKeys() []Key {
	return m.mandatoryKeys
}

func (m *ManagerImpl) SetMandatoryKeys(mandatoryKeys []Key) error {
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

func validateConfiguration(config Config, mandatoryKeys []Key) error {
	missingKey := ""
	containsMandatoryKeys := lo.ContainsBy(mandatoryKeys, func(key Key) bool {
		containsKey := lo.ContainsBy(config.Keys, func(item core.ConfigurationKey) bool {
			return item.Key == key.String()
		})

		if !containsKey {
			missingKey = key.String()
		}

		return containsKey
	})

	if !containsMandatoryKeys {
		return fmt.Errorf("missing mandatory keys: %s", missingKey)
	}

	return nil
}

func (m *ManagerImpl) UpdateKey(key Key, value *string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.keyValidator != nil {
		if !m.keyValidator(key, value) {
			return fmt.Errorf("key validation failed for key %s", key)
		}
	}

	return m.ocppConfig.UpdateKey(key.String(), value)
}

func (m *ManagerImpl) GetConfiguration() ([]core.ConfigurationKey, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfig(), nil
}

func (m *ManagerImpl) GetConfigurationValue(key Key) (*string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfigurationValue(key.String())
}
