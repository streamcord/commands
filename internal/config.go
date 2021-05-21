package internal

import (
	log "github.com/sirupsen/logrus"
	"github.com/streamcord/commands/types"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// GlobalConf - Global instance of the application config.
var GlobalConf types.Config

// LoadConfig reads a YAML file and parse it into the application config.
func LoadConfig(path string) {

	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("Failed to read yaml file", err)
	}

	err = yaml.Unmarshal(file, &GlobalConf)
	if err != nil {
		log.Fatal("Failed to parse yaml", err)
	}
}
