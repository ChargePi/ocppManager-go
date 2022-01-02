package manager

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

const (
	defaultFileFormat = YamlFile
	defaultFileName   = "configuration"
	defaultFilePath   = "."
)

type ConfigurationManagerTestSuite struct {
	suite.Suite
	config  configuration.Config
	manager Manager
}

func (s *ConfigurationManagerTestSuite) SetupTest() {
	s.manager = NewManager()

	// Set default file information
	s.manager.SetFilePath(defaultFilePath)
	s.manager.SetFileName(defaultFileName)
	s.manager.SetFileFormat(defaultFileFormat)

	// Set supported profile and version
	s.manager.SetSupportedProfiles("core")
	s.manager.SetVersion(configuration.OCPP16)

	s.config = configuration.Config{
		Version: 1,
		Keys: []core.ConfigurationKey{
			{
				Key:      "Test1",
				Readonly: false,
				Value:    "60",
			}, {
				Key:      "Test2",
				Readonly: false,
				Value:    "ABCD",
			},
			{
				Key:      "MeterValuesSampledData",
				Readonly: false,
				Value: strings.Join(
					[]string{
						string(types.MeasurandEnergyActiveExportInterval),
						string(types.MeasurandCurrentExport),
						string(types.MeasurandVoltage),
					},
					",",
				),
			},
		}}

	err := writeToFile("configuration.yaml", s.config)
	s.Require().NoError(err)
}

func (s *ConfigurationManagerTestSuite) TestGetConfiguration() {
	err := s.manager.LoadConfiguration()
	s.Require().NoError(err)

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
	s.Require().NoError(err)

	err = fig.Load(&fileConfig, fig.File("configuration.yaml"), fig.Dirs("."))
	s.Require().NoError(err)
	s.Require().Equal(s.config, fileConfig)

	// Delete the unnecessary file
	exec.Command("rm configuration.json")
}

func TestConfigurationManager(t *testing.T) {
	suite.Run(t, new(ConfigurationManagerTestSuite))
}
