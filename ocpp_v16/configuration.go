package ocpp_v16

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/firmware"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
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
			Value:    lo.ToPtr("0"),
		},
		{
			Key:      ConnectionTimeOut.String(),
			Readonly: false,
			Value:    lo.ToPtr("60"),
		},
		{
			Key:      GetConfigurationMaxKeys.String(),
			Readonly: false,
			Value:    lo.ToPtr("100"),
		},
		{
			Key:      HeartbeatInterval.String(),
			Readonly: false,
			Value:    lo.ToPtr("60"),
		},
		{
			Key:      LocalPreAuthorize.String(),
			Readonly: false,
			Value:    lo.ToPtr("false"),
		},
		{
			Key:      MeterValuesAlignedData.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
		{
			Key:      MeterValuesSampledData.String(),
			Readonly: false,
			Value: lo.ToPtr(strings.Join([]string{
				string(types.MeasurandVoltage),
				string(types.MeasurandCurrentImport),
				string(types.MeasurandPowerActiveImport),
				string(types.MeasurandEnergyActiveImportInterval),
				string(types.MeasueandSoC),
			}, ",")),
		},
		{
			Key:      MeterValueSampleInterval.String(),
			Readonly: false,
			Value:    lo.ToPtr("20"),
		},
		{
			Key:      NumberOfConnectors.String(),
			Readonly: false,
			Value:    lo.ToPtr("1"),
		},
		{
			Key:      ResetRetries.String(),
			Readonly: false,
			Value:    lo.ToPtr("3"),
		},
		{
			Key:      ConnectorPhaseRotation.String(),
			Readonly: false,
			Value:    lo.ToPtr("Unknown"),
		},
		{
			Key:      StopTransactionOnEVSideDisconnect.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
		{
			Key:      StopTransactionOnInvalidId.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
		{
			Key:      StopTxnAlignedData.String(),
			Readonly: false,
			Value: lo.ToPtr(strings.Join([]string{
				string(types.MeasurandVoltage),
				string(types.MeasurandCurrentImport),
				string(types.MeasurandPowerActiveImport),
				string(types.MeasurandEnergyActiveImportInterval),
				string(types.MeasueandSoC),
			}, ",")),
		},
		{
			Key:      StopTxnSampledData.String(),
			Readonly: false,
			Value: lo.ToPtr(strings.Join([]string{
				string(types.MeasurandVoltage),
				string(types.MeasurandCurrentImport),
				string(types.MeasurandPowerActiveImport),
				string(types.MeasurandEnergyActiveImportInterval),
				string(types.MeasueandSoC),
			}, ",")),
		},
		{
			Key:      SupportedFeatureProfiles.String(),
			Readonly: true,
			Value:    lo.ToPtr("Core"),
		},
		{
			Key:      TransactionMessageAttempts.String(),
			Readonly: false,
			Value:    lo.ToPtr("3"),
		},
		{
			Key:      TransactionMessageRetryInterval.String(),
			Readonly: false,
			Value:    lo.ToPtr("30"),
		},
		{
			Key:      UnlockConnectorOnEVSideDisconnect.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
	}
}

func DefaultLocalAuthConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      LocalAuthListEnabled.String(),
			Readonly: false,
			Value:    lo.ToPtr("true"),
		},
		{
			Key:      LocalAuthListMaxLength.String(),
			Readonly: true,
			Value:    lo.ToPtr("10"),
		},
		{
			Key:      SendLocalListMaxLength.String(),
			Readonly: true,
			Value:    lo.ToPtr("10"),
		},
	}
}

func DefaultSmartChargingConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      ChargeProfileMaxStackLevel.String(),
			Readonly: true,
			Value:    lo.ToPtr("5"),
		},
		{
			Key:      ChargingScheduleAllowedChargingRateUnit.String(),
			Readonly: true,
			Value:    lo.ToPtr("Current,Power"),
		},
		{
			Key:      ChargingScheduleMaxPeriods.String(),
			Readonly: true,
			Value:    lo.ToPtr("6"),
		},
		{
			Key:      MaxChargingProfilesInstalled.String(),
			Readonly: true,
			Value:    lo.ToPtr("5"),
		},
	}
}

func DefaultFirmwareConfiguration() []core.ConfigurationKey {
	return []core.ConfigurationKey{
		{
			Key:      SupportedFileTransferProtocols.String(),
			Readonly: true,
			Value:    lo.ToPtr("HTTP,HTTPS,FTP,FTPS,SFTP"),
		},
	}
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
