package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf = new(config)

// MySQLConfig 数据库配置

type config struct {
	AppConfig struct {
		Release bool `yaml:"release"`
		Port    int  `yaml:"port"`
	} `yaml:"appConfig"`

	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		DB       string `yaml:"db"`
	} `yaml:"mysql"`
	Log struct {
		Filename   string
		MaxSize    int
		MaxBackups int
		MaxAge     int
		Compress   bool
	}
	JwtSecret string `yaml:"jwt_secret"`
}

func init() {
	yamlFile, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Conf)
	fmt.Println("Conf ", Conf)
}
