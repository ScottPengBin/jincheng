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
}

type Config struct {
	MySQLConf `yaml:"mysql"`
}

var Provider = wire.NewSet(GetConfig)

func GetConfig() Config {
	b, _ := ioutil.ReadFile("config.yml")

	var conf Config
	_ = yaml.Unmarshal(b, &conf)

	return conf
}
