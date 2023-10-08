package ocpp_v16

import (
	"errors"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
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

func NewEmptyConfiguration() Config {
	return Config{
		Version: 1,
		Keys:    []core.ConfigurationKey{},
	}
}

func DefaultConfiguration(profiles ...string) Config {
	keys := DefaultCoreConfiguration()

	for _, profile := range profiles {
		switch profile {
		case localauth.ProfileName:
			keys = append(keys, DefaultLocalAuthConfiguration()...)
		case smartcharging.ProfileName:
			keys = append(keys, DefaultSmartChargingConfiguration()...)
		case firmware.ProfileName:
			keys = append(keys, DefaultFirmwareConfiguration()...)
		}
	}

	return Config{
		Version: 1,
		Keys:    keys,
	}
}

func DefaultCoreConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      AuthorizeRemoteTxRequests.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
		{
			Key:      ClockAlignedDataInterval.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ConnectionTimeOut.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      GetConfigurationMaxKeys.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      HeartbeatInterval.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      LocalPreAuthorize.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      MeterValuesAlignedData.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      MeterValuesSampledData.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      MeterValuesSampledData.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      MeterValueSampleInterval.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      NumberOfConnectors.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ResetRetries.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ConnectorPhaseRotation.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      StopTransactionOnEVSideDisconnect.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      StopTransactionOnInvalidId.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      StopTxnAlignedData.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      StopTxnSampledData.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      SupportedFeatureProfiles.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      TransactionMessageAttempts.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      TransactionMessageRetryInterval.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      UnlockConnectorOnEVSideDisconnect.String(),
			Readonly: false,
			Value:    nil,
		},
	}
}

func DefaultLocalAuthConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      LocalAuthListEnabled.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      LocalAuthListMaxLength.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      SendLocalListMaxLength.String(),
			Readonly: false,
			Value:    nil,
		},
	}
}

func DefaultSmartChargingConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      ChargeProfileMaxStackLevel.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ChargingScheduleAllowedChargingRateUnit.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ChargingScheduleMaxPeriods.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      MaxChargingProfilesInstalled.String(),
			Readonly: false,
			Value:    nil,
		},
		{
			Key:      ConnectorSwitch3to1PhaseSupported.String(),
			Readonly: false,
			Value:    nil,
		},
	}
}

func DefaultFirmwareConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      SupportedFileTransferProtocols.String(),
			Readonly: false,
			Value:    nil,
		},
	}
}

// UpdateKey Update the configuration variable in the configuration if it is not readonly.
func (config *Config) UpdateKey(key string, value *string) error {
	log.Debugf("Updating key %s", key)

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
