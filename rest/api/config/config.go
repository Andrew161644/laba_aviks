package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Host            string
	Port            string
	DbHost          string
	DbPort          int
	DbUsername      string
	DbName          string
	DbPassword      string
	CalcUri         string
	RabbitHost      string
	LocalRabbitHost string
	QueueName       string
}

func GetConfig() (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("../resources/conf.toml", &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
