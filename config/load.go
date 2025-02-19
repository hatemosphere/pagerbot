package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

var Config config

func Load(filePath string) error {
	Config = config{}
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return fmt.Errorf("config file not found")
	}

	configContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal([]byte(os.ExpandEnv(string(configContent))), &Config)
	if err != nil {
		return fmt.Errorf("error parsing config file: %s", err)
	}

	return nil
}
