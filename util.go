package ocppConfigManager

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strings"
)

const (
	YamlFile = FileFormat("yaml")
	YmlFile  = FileFormat("yml")
	JSON     = FileFormat("json")
	TOML     = FileFormat("toml")
)

type FileFormat string

var (
	ErrUnsupportedFileFormat = errors.New("unsupported file format")
)

// writeToFile writes come JSON/YAML/TOML structure to the specified path.
func writeToFile(filename string, structure interface{}) (err error) {
	log.Debugf("Creating a file: %s", filename)

	var (
		encodingType          string
		marshal               []byte
		splitFile             = strings.Split(filename, ".")
		isValidFile, matchErr = regexp.MatchString("^.*\\.(json|yaml|yml|toml)$", filename)
	)

	if matchErr != nil {
		return matchErr
	}

	// Check if the file format is supported
	if len(splitFile) > 0 && isValidFile {
		encodingType = splitFile[len(splitFile)-1]
	}

	switch FileFormat(encodingType) {
	case YamlFile, YmlFile:
		marshal, err = yaml.Marshal(&structure)
		break
	case JSON:
		marshal, err = json.MarshalIndent(&structure, "", "\t")
		break
	default:
		return ErrUnsupportedFileFormat
	}

	return ioutil.WriteFile(filename, marshal, 0644)
}
