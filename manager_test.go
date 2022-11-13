package ocppConfigManager

import (
	"github.com/kkyr/fig"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/core"
	"github.com/lorenzodonini/ocpp-go/ocpp1.6/types"
	"github.com/stretchr/testify/suite"
	"github.com/xBlaz3kx/ocppManager-go/configuration"
	"os/exec"
	"strings"
	"testing"
)

type ConfigurationManagerTestSuite struct {
	suite.Suite
	config  configuration.Config
	manager Manager
}

func (s *ConfigurationManagerTestSuite) SetupTest() {
	s.manager = NewManager()

	// Set default file information
	s.manager.SetFilePath(defaultFilePath + defaultFileName + string(defaultFileFormat))

	// Set supported profile and version
	s.manager.SetSupportedProfiles("core")
	s.manager.SetVersion(configuration.OCPP16)

	val := strings.Join(
		[]string{
			string(types.MeasurandEnergyActiveExportInterval),
			string(types.MeasurandCurrentExport),
			string(types.MeasurandVoltage),
		},
		",")

	val1 := "60"

	s.config = configuration.Config{
		Version: 1,
		Keys: []core.ConfigurationKey{
			{
				Key:      configuration.HeartbeatInterval.String(),
				Readonly: false,
				Value:    &val1,
			}, {
				Key:      configuration.ChargeProfileMaxStackLevel.String(),
				Readonly: false,
				Value:    nil,
			},
			{
				Key:      configuration.MeterValuesSampledData.String(),
				Readonly: false,
				Value:    &val,
			},
		}}

	err := writeToFile("configuration.yaml", s.config)
	s.Assert().NoError(err)
}

func (s *ConfigurationManagerTestSuite) TestGetConfiguration() {
	err := s.manager.LoadConfiguration()
	s.Assert().NoError(err)

	//todo
}

func (s *ConfigurationManagerTestSuite) TestUpdateConfiguration() {
	//todo
}

func (s *ConfigurationManagerTestSuite) TestGetConfigurationValue() {
	//todo
}

func (s *ConfigurationManagerTestSuite) TestUpdateConfigurationFile() {
	var (
		fileConfig configuration.Config
	)

	err := s.manager.UpdateConfigurationFile()
	s.Assert().NoError(err)

	err = fig.Load(&fileConfig, fig.File("configuration.yaml"), fig.Dirs("."))
	s.Assert().NoError(err)
	s.Assert().Equal(s.config, fileConfig)

	// Delete the unnecessary file
	exec.Command("rm configuration.json")
}

func TestConfigurationManager(t *testing.T) {
	suite.Run(t, new(ConfigurationManagerTestSuite))
}
