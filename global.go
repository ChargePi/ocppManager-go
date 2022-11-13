package ocppConfigManager

import (
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/xBlaz3kx/ocppManager-go/configuration"
	"reflect"
	"sync"
)

const (
	defaultFileFormat = YamlFile
	defaultFileName   = "configuration"
	defaultFilePath   = "."
)

var (
	manager Manager
	once    = sync.Once{}
)

func init() {
	once.Do(func() {
		manager = NewManager()
		// Set default file information
		manager.SetFilePath(defaultFilePath)
		manager.SetFilePath(defaultFileName)

		// Set supported profile and version
		manager.SetSupportedProfiles("core")
		manager.SetVersion(configuration.OCPP16)
	})
}

func isNilInterfaceOrPointer(sth interface{}) bool {
	return sth == nil || (reflect.ValueOf(sth).Kind() == reflect.Ptr && reflect.ValueOf(sth).IsNil())
}

func GetManager() Manager {
	return manager
}

func SetManager(confManager Manager) {
	if isNilInterfaceOrPointer(confManager) {
		return
	}

	manager = confManager
}

func SetFilePath(path string) {
	manager.SetFilePath(path)
}

func SetSupportedProfiles(profiles ...string) {
	manager.SetSupportedProfiles(profiles...)
}

func SetVersion(version configuration.ProtocolVersion) {
	manager.SetVersion(version)
}

// LoadConfiguration load OCPP configuration from the file.
func LoadConfiguration() error {
	return manager.LoadConfiguration()
}

// GetConfiguration Get the global configuration
func GetConfiguration() ([]core.ConfigurationKey, error) {
	return manager.GetConfiguration()
}

// UpdateKey Update the configuration variable in the global configuration if it is not readonly.
func UpdateKey(key string, value *string) error {
	return manager.UpdateKey(key, value)
}

// GetConfigurationValue Get the value of specified configuration variable from the global configuration in String format.
func GetConfigurationValue(key string) (*string, error) {
	return manager.GetConfigurationValue(key)
}

// UpdateConfigurationFile Write/Rewrite the configuration to the file.
func UpdateConfigurationFile() error {
	return manager.UpdateConfigurationFile()
}
