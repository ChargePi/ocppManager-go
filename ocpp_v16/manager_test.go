package ocpp_v16

import (
	"strings"
	"testing"

	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/localauth"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/smartcharging"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
	"github.com/stretchr/testify/suite"
)

type ConfigurationManagerTestSuite struct {
	suite.Suite
	config  Config
	manager Manager
}

func (s *ConfigurationManagerTestSuite) SetupTest() {
	val := strings.Join(
		[]string{
			string(types.MeasurandEnergyActiveExportInterval),
			string(types.MeasurandCurrentExport),
			string(types.MeasurandVoltage),
		},
		",")

	val1 := "60"

	s.config = Config{
		Version: 1,
		Keys: []core.ConfigurationKey{
			{
				Key:      HeartbeatInterval.String(),
				Readonly: false,
				Value:    &val1,
			}, {
				Key:      ChargeProfileMaxStackLevel.String(),
				Readonly: false,
				Value:    nil,
			},
			{
				Key:      MeterValuesSampledData.String(),
				Readonly: false,
				Value:    &val,
			},
		}}

	var err error
	s.manager, err = NewV16ConfigurationManager(s.config, core.ProfileName, smartcharging.ProfileName, localauth.ProfileName)

	s.Assert().NoError(err)
}

func (s *ConfigurationManagerTestSuite) TestGetConfiguration() {
	// todo
}

func (s *ConfigurationManagerTestSuite) TestUpdateConfiguration() {
	// todo
}

func (s *ConfigurationManagerTestSuite) TestGetConfigurationValue() {
	// todo
}

func TestConfigurationManager(t *testing.T) {
	suite.Run(t, new(ConfigurationManagerTestSuite))
}
