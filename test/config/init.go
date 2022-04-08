package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf = new(config)

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
}
