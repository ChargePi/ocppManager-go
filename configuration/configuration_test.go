package configuration

import (
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

var (
	val1 = "60"
	val2 = "ABCD"
)

type OcppConfigTest struct {
	suite.Suite
	keys   []core.ConfigurationKey
	config Config
}

func (s *OcppConfigTest) SetupTest() {
	s.keys = []core.ConfigurationKey{
		{
			Key:      HeartbeatInterval.String(),
			Readonly: false,
			Value:    &val1,
		}, {
			Key:      ChargingScheduleAllowedChargingRateUnit.String(),
			Readonly: true,
			Value:    &val2,
		}, {
			Key:      AuthorizationCacheEnabled.String(),
			Readonly: false,
			Value:    nil,
		},
	}

	s.config = Config{
		Version: 1,
		Keys:    s.keys,
	}
}

func (s *OcppConfigTest) TestGetConfig() {
	s.Assert().Equal(s.keys, s.config.GetConfig())

	// Overwrite the config
	s.config = Config{
		Version: 1,
		Keys:    []core.ConfigurationKey{},
	}

	s.Assert().Equal([]core.ConfigurationKey{}, s.config.GetConfig())
}

func (s *OcppConfigTest) TestUpdateKey() {
	// Ok casew
	newVal := "1234"
	err := s.config.UpdateKey(HeartbeatInterval.String(), &newVal)
	s.Assert().NoError(err)
	value, err := s.config.GetConfigurationValue(HeartbeatInterval.String())
	s.Require().NoError(err)
	s.Assert().EqualValues("1234", *value)

	// Invalid key
	err = s.config.UpdateKey("Test4", nil)
	s.Assert().Error(err)

	// Key is readOnly
	err = s.config.UpdateKey(ChargingScheduleAllowedChargingRateUnit.String(), nil)
	s.Assert().Error(err)
	value, err = s.config.GetConfigurationValue(ChargingScheduleAllowedChargingRateUnit.String())
	s.Assert().NoError(err)
	s.Assert().EqualValues("ABCD", *value)
}

func TestOCPPConfig(t *testing.T) {
	suite.Run(t, new(OcppConfigTest))
}
