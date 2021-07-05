package internal

import (
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/streamcord/commands/types/config"
	"gopkg.in/yaml.v2"
)

// GlobalConf - Global instance of the application config.
var GlobalConf config.Config

// LoadConfig reads a YAML file and parse it into the application config.
func LoadConfig(path string) config.Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read yaml file: %s", err.Error())
	}

	var conf config.Config

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("Failed to parse yaml: %s", err.Error())
	}

	return conf
}
