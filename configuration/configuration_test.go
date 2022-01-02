package configuration

import (
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OcppConfigTest struct {
	suite.Suite
	keys   []core.ConfigurationKey
	config Config
}

func (s *OcppConfigTest) SetupTest() {
	s.keys = []core.ConfigurationKey{
		{
			Key:      "Test1",
			Readonly: false,
			Value:    "60",
		}, {
			Key:      "Test2",
			Readonly: false,
			Value:    "ABCD",
		}, {
			Key:      "Test3",
			Readonly: false,
			Value:    "",
		},
	}

	s.config = Config{
		Version: 1,
		Keys:    s.keys,
	}
}

func (s *OcppConfigTest) TestGetConfig() {
	s.Require().Equal(s.keys, s.config.GetConfig())

	// Overwrite the config
	s.config = Config{
		Version: 1,
		Keys:    []core.ConfigurationKey{},
	}

	s.Require().Equal([]core.ConfigurationKey{}, s.config.GetConfig())
}

func (s *OcppConfigTest) TestUpdateKey() {
	// Ok case
	err := s.config.UpdateKey("Test1", "1234")
	s.Require().NoError(err)
	value, err := s.config.GetConfigurationValue("Test1")
	s.Require().NoError(err)
	s.Require().Equal("1234", value)

	// Invalid key
	err = s.config.UpdateKey("Test4", "1234")
	s.Require().Error(err)

	// Key is readOnly
	err = s.config.UpdateKey("Test2", "BCDEF")
	s.Require().Error(err)
	value, err = s.config.GetConfigurationValue("Test2")
	s.Require().NoError(err)
	s.Require().Equal("ABCD", value)
}

func TestOCPPConfig(t *testing.T) {
	suite.Run(t, new(OcppConfigTest))
}
