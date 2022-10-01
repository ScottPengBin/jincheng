package config

import (
	"github.com/google/wire"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type MySQLConf struct {
	Driver string `yaml:"driver"`
	Master struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"master"`
	Slave struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"slave"`
	Prefix string `yaml:"prefix"`
}

type Config struct {
	MySQLConf `yaml:"mysql"`
	App       `yaml:"app"`
}

type App struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

var Provider = wire.NewSet(GetConfig)

func GetConfig() *Config {
	b, _ := ioutil.ReadFile("config.yml")

	var conf Config
	_ = yaml.Unmarshal(b, &conf)

	return &conf
}
