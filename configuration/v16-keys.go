package configuration

const (
	/* ----------------- Core keys ----------------------- */

	AllowOfflineTxForUnknownId        = Key("AllowOfflineTxForUnknownId")
	AuthorizationCacheEnabled         = Key("AuthorizationCacheEnabled")
	AuthorizeRemoteTxRequests         = Key("AuthorizeRemoteTxRequests")
	BlinkRepeat                       = Key("BlinkRepeat")
	ClockAlignedDataInterval          = Key("ClockAlignedDataInterval")
	ConnectionTimeOut                 = Key("ConnectionTimeOut")
	GetConfigurationMaxKeys           = Key("GetConfigurationMaxKeys")
	HeartbeatInterval                 = Key("HeartbeatInterval")
	LightIntensity                    = Key("LightIntensity")
	LocalAuthorizeOffline             = Key("LocalAuthorizeOffline")
	LocalPreAuthorize                 = Key("LocalPreAuthorize")
	MaxEnergyOnInvalidId              = Key("MaxEnergyOnInvalidId")
	MeterValuesAlignedData            = Key("MeterValuesAlignedData")
	MeterValuesAlignedDataMaxLength   = Key("MeterValuesAlignedDataMaxLength")
	MeterValuesSampledData            = Key("MeterValuesSampledData")
	MeterValuesSampledDataMaxLength   = Key("MeterValuesSampledDataMaxLength")
	MeterValueSampleInterval          = Key("MeterValueSampleInterval")
	MinimumStatusDuration             = Key("MinimumStatusDuration")
	NumberOfConnectors                = Key("NumberOfConnectors")
	ResetRetries                      = Key("ResetRetries")
	ConnectorPhaseRotation            = Key("ConnectorPhaseRotation")
	ConnectorPhaseRotationMaxLength   = Key("ConnectorPhaseRotationMaxLength")
	StopTransactionOnEVSideDisconnect = Key("StopTransactionOnEVSideDisconnect")
	StopTransactionOnInvalidId        = Key("StopTransactionOnInvalidId")
	StopTxnAlignedData                = Key("StopTxnAlignedData")
	StopTxnAlignedDataMaxLength       = Key("StopTxnAlignedDataMaxLength")
	StopTxnSampledData                = Key("StopTxnSampledData")
	StopTxnSampledDataMaxLength       = Key("StopTxnSampledDataMaxLength")
	SupportedFeatureProfiles          = Key("SupportedFeatureProfiles")
	SupportedFeatureProfilesMaxLength = Key("SupportedFeatureProfilesMaxLength")
	TransactionMessageAttempts        = Key("TransactionMessageAttempts")
	TransactionMessageRetryInterval   = Key("TransactionMessageRetryInterval")
	UnlockConnectorOnEVSideDisconnect = Key("UnlockConnectorOnEVSideDisconnect")
	WebSocketPingInterval             = Key("WebSocketPingInterval")

	/* ----------------- LocalAuthList keys ----------------------- */

	LocalAuthListEnabled   = Key("LocalAuthListEnabled")
	LocalAuthListMaxLength = Key("LocalAuthListMaxLength")
	SendLocalListMaxLength = Key("SendLocalListMaxLength")

	/* ----------------- Reservation keys ----------------------- */

	ReserveConnectorZeroSupported = Key("ReserveConnectorZeroSupported")

	/* ----------------- SmartCharging keys ----------------------- */

	ChargeProfileMaxStackLevel              = Key("ChargeProfileMaxStackLevel")
	ChargingScheduleAllowedChargingRateUnit = Key("ChargingScheduleAllowedChargingRateUnit")
	ChargingScheduleMaxPeriods              = Key("ChargingScheduleMaxPeriods")
	MaxChargingProfilesInstalled            = Key("MaxChargingProfilesInstalled")
	ConnectorSwitch3to1PhaseSupported       = Key("ConnectorSwitch3to1PhaseSupported")
)

var (
	MandatoryCoreKeys = []Key{
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

	MandatoryLocalAuthKeys = []Key{
		LocalAuthListEnabled,
		LocalAuthListMaxLength,
		SendLocalListMaxLength,
	}

	MandatorySmartChargingKeys = []Key{
		MaxChargingProfilesInstalled,
		ChargingScheduleMaxPeriods,
		ChargingScheduleAllowedChargingRateUnit,
		ChargeProfileMaxStackLevel,
	}
)
