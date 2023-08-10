package config

import (
	"niltasks/pkg/mongo/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mongo config.MongoConfig `yaml:"mongo"`
}

func MustLoad() *Config {
	var cfg Config

	err := cleanenv.ReadConfig("config/config.yml", &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
