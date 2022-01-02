package manager

import (
	"fmt"
	"github.com/kkyr/fig"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	log "github.com/sirupsen/logrus"
	"github.com/xBlaz3kx/ocppManager-go/configuration"
	v16 "github.com/xBlaz3kx/ocppManager-go/v16"
	"path/filepath"
	"strings"
	"sync"
)

type (
	Manager interface {
		SetFileFormat(format FileFormat)
		SetFileName(name string)
		SetFilePath(path string)
		SetVersion(version configuration.ProtocolVersion)
		SetSupportedProfiles(profiles ...string)
		SetConfiguration(configuration configuration.Config) error
		LoadConfiguration() error
		validateConfiguration(config *configuration.Config, mandatoryKeys []configuration.Key) error
		UpdateKey(key string, value string) error
		GetConfigurationValue(key string) (string, error)
		GetConfiguration() ([]core.ConfigurationKey, error)
		UpdateConfigurationFile() error
	}

	ManagerImpl struct {
		fileFormat        FileFormat
		fileName          string
		filePath          string
		fullFilePath      string
		supportedProfiles []string
		ocppConfig        *configuration.Config
		ocppVersion       configuration.ProtocolVersion
		mandatoryKeys     []configuration.Key
		mu                sync.Mutex
	}
)

func NewManager() Manager {
	return &ManagerImpl{
		fileFormat:    "",
		fileName:      "",
		filePath:      "",
		fullFilePath:  "",
		ocppConfig:    nil,
		ocppVersion:   "",
		mandatoryKeys: []configuration.Key{},
		mu:            sync.Mutex{},
	}
}

func (m *ManagerImpl) SetFileFormat(format FileFormat) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.fileFormat = format
	m.fullFilePath = fmt.Sprintf("%s/%s.%s", m.filePath, m.fileName, m.fileFormat)
}

func (m *ManagerImpl) SetFileName(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.fileName = name
	m.fullFilePath = fmt.Sprintf("%s/%s.%s", m.filePath, m.fileName, m.fileFormat)
}

func (m *ManagerImpl) SetFilePath(path string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.filePath = path
	m.fullFilePath = fmt.Sprintf("%s/%s.%s", m.filePath, m.fileName, m.fileFormat)
}

func (m *ManagerImpl) SetSupportedProfiles(profiles ...string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.supportedProfiles = profiles
	m.fullFilePath = fmt.Sprintf("%s/%s.%s", m.filePath, m.fileName, m.fileFormat)
}

func (m *ManagerImpl) SetConfiguration(configuration configuration.Config) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	err := m.validateConfiguration(&configuration, m.mandatoryKeys)
	if err != nil {
		return err
	}

	m.ocppConfig = &configuration
	return nil
}

func (m *ManagerImpl) SetVersion(version configuration.ProtocolVersion) {
	m.mu.Lock()
	defer m.mu.Unlock()

	switch version {
	case configuration.OCPP16:
		for _, profile := range m.supportedProfiles {
			switch strings.ToLower(profile) {
			case core.ProfileName:
				m.mandatoryKeys = append(m.mandatoryKeys, v16.MandatoryCoreKeys...)
				break
			case smartcharging.ProfileName:
				m.mandatoryKeys = append(m.mandatoryKeys, v16.MandatorySmartChargingKeys...)
				break
			case localauth.ProfileName:
				m.mandatoryKeys = append(m.mandatoryKeys, v16.MandatoryLocalAuthKeys...)
				break
			}
		}
	case configuration.OCPP201:
		// todo when available
		break
	default:
		return
	}

	m.ocppVersion = version
}

func (m *ManagerImpl) LoadConfiguration() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.mandatoryKeys == nil {
		log.Fatalf("No mandatory keys applied")
	}

	config := configuration.Config{}
	err := fig.Load(&config,
		fig.File(filepath.Base(m.fullFilePath)),
		fig.Dirs(filepath.Dir(m.fullFilePath)),
	)
	if err != nil {
		log.Fatalf("Error loading file: %v", err)
		return err
	}

	m.ocppConfig = &config
	return m.validateConfiguration(m.ocppConfig, m.mandatoryKeys)
}

func (m *ManagerImpl) validateConfiguration(config *configuration.Config, mandatoryKeys []configuration.Key) error {
	containsMandatoryKeys := true
	missingKey := ""

	for _, mandatoryKey := range mandatoryKeys {
		hasKey := false

		for _, key := range config.Keys {
			if mandatoryKey.String() == key.Key {
				hasKey = true
				break
			}
		}

		if !hasKey {
			containsMandatoryKeys = false
			missingKey = mandatoryKey.String()
			break
		}
	}

	if !containsMandatoryKeys {
		return fmt.Errorf("missing mandatory keys: %s", missingKey)
	}

	return nil
}

func (m *ManagerImpl) UpdateKey(key string, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.UpdateKey(key, value)
}

func (m *ManagerImpl) GetConfiguration() ([]core.ConfigurationKey, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfig(), nil
}

func (m *ManagerImpl) GetConfigurationValue(key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.ocppConfig.GetConfigurationValue(key)
}

func (m *ManagerImpl) UpdateConfigurationFile() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return writeToFile(m.fullFilePath, m.ocppConfig)
}
