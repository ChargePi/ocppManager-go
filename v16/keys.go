package v16

import (
	"gihub.com/xBlaz3kx/ocppManager-go/configuration"
)

const (
	/* ----------------- Core keys ----------------------- */

	AllowOfflineTxForUnknownId        = configuration.Key("AllowOfflineTxForUnknownId")
	AuthorizationCacheEnabled         = configuration.Key("AuthorizationCacheEnabled")
	AuthorizeRemoteTxRequests         = configuration.Key("AuthorizeRemoteTxRequests")
	BlinkRepeat                       = configuration.Key("BlinkRepeat")
	ClockAlignedDataInterval          = configuration.Key("ClockAlignedDataInterval")
	ConnectionTimeOut                 = configuration.Key("ConnectionTimeOut")
	GetConfigurationMaxKeys           = configuration.Key("GetConfigurationMaxKeys")
	HeartbeatInterval                 = configuration.Key("HeartbeatInterval")
	LightIntensity                    = configuration.Key("LightIntensity")
	LocalAuthorizeOffline             = configuration.Key("LocalAuthorizeOffline")
	LocalPreAuthorize                 = configuration.Key("LocalPreAuthorize")
	MaxEnergyOnInvalidId              = configuration.Key("MaxEnergyOnInvalidId")
	MeterValuesAlignedData            = configuration.Key("MeterValuesAlignedData")
	MeterValuesAlignedDataMaxLength   = configuration.Key("MeterValuesAlignedDataMaxLength")
	MeterValuesSampledData            = configuration.Key("MeterValuesSampledData")
	MeterValuesSampledDataMaxLength   = configuration.Key("MeterValuesSampledDataMaxLength")
	MeterValueSampleInterval          = configuration.Key("MeterValueSampleInterval")
	MinimumStatusDuration             = configuration.Key("MinimumStatusDuration")
	NumberOfConnectors                = configuration.Key("NumberOfConnectors")
	ResetRetries                      = configuration.Key("ResetRetries")
	ConnectorPhaseRotation            = configuration.Key("ConnectorPhaseRotation")
	ConnectorPhaseRotationMaxLength   = configuration.Key("ConnectorPhaseRotationMaxLength")
	StopTransactionOnEVSideDisconnect = configuration.Key("StopTransactionOnEVSideDisconnect")
	StopTransactionOnInvalidId        = configuration.Key("StopTransactionOnInvalidId")
	StopTxnAlignedData                = configuration.Key("StopTxnAlignedData")
	StopTxnAlignedDataMaxLength       = configuration.Key("StopTxnAlignedDataMaxLength")
	StopTxnSampledData                = configuration.Key("StopTxnSampledData")
	StopTxnSampledDataMaxLength       = configuration.Key("StopTxnSampledDataMaxLength")
	SupportedFeatureProfiles          = configuration.Key("SupportedFeatureProfiles")
	SupportedFeatureProfilesMaxLength = configuration.Key("SupportedFeatureProfilesMaxLength")
	TransactionMessageAttempts        = configuration.Key("TransactionMessageAttempts")
	TransactionMessageRetryInterval   = configuration.Key("TransactionMessageRetryInterval")
	UnlockConnectorOnEVSideDisconnect = configuration.Key("UnlockConnectorOnEVSideDisconnect")
	WebSocketPingInterval             = configuration.Key("WebSocketPingInterval")

	/* ----------------- LocalAuthList keys ----------------------- */

	LocalAuthListEnabled   = configuration.Key("LocalAuthListEnabled")
	LocalAuthListMaxLength = configuration.Key("LocalAuthListMaxLength")
	SendLocalListMaxLength = configuration.Key("SendLocalListMaxLength")

	/* ----------------- Reservation keys ----------------------- */

	ReserveConnectorZeroSupported = configuration.Key("ReserveConnectorZeroSupported")

	/* ----------------- SmartCharging keys ----------------------- */

	ChargeProfileMaxStackLevel              = configuration.Key("ChargeProfileMaxStackLevel")
	ChargingScheduleAllowedChargingRateUnit = configuration.Key("ChargingScheduleAllowedChargingRateUnit")
	ChargingScheduleMaxPeriods              = configuration.Key("ChargingScheduleMaxPeriods")
	MaxChargingProfilesInstalled            = configuration.Key("MaxChargingProfilesInstalled")
	ConnectorSwitch3to1PhaseSupported       = configuration.Key("ConnectorSwitch3to1PhaseSupported")
)

var (
	MandatoryCoreKeys = []configuration.Key{
		AuthorizeRemoteTxRequests,
		ClockAlignedDataInterval,
		ConnectionTimeOut,
		GetConfigurationMaxKeys,
		HeartbeatInterval,
		LocalPreAuthorize,
		MeterValuesAlignedData,
		MeterValuesSampledData,
		MeterValueSampleInterval,
		NumberOfConnectors,
		ResetRetries,
		ConnectorPhaseRotation,
		StopTransactionOnEVSideDisconnect,
		StopTransactionOnInvalidId,
		StopTxnAlignedData,
		StopTxnSampledData,
		SupportedFeatureProfiles,
		TransactionMessageAttempts,
		TransactionMessageRetryInterval,
		UnlockConnectorOnEVSideDisconnect,
	}

	MandatoryLocalAuthKeys = []configuration.Key{
		LocalAuthListEnabled,
		LocalAuthListMaxLength,
		SendLocalListMaxLength,
	}

	MandatorySmartChargingKeys = []configuration.Key{
		MaxChargingProfilesInstalled,
		ChargingScheduleMaxPeriods,
		ChargingScheduleAllowedChargingRateUnit,
		ChargeProfileMaxStackLevel,
	}
)
