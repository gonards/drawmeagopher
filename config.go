package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

const configFile = "config.yml"

// Conf loaded
var Conf *Config

func parseConfig() {
	if Conf == nil {
		file, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(err)
		}
		Conf = new(Config)
		err = yaml.Unmarshal(file, &Conf)
		if err != nil {
			panic(err)
		}
	}
}
