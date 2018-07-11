package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const settingsFile = "settings.yml"

// Conf loaded
var settings *Settings

func parseSettings() error {
	file, err := ioutil.ReadFile(settingsFile)
	if err != nil {
		return err
	}
	settings = new(Settings)
	err = yaml.Unmarshal(file, &settings)
	if err != nil {
		return err
	}
	return nil
}
