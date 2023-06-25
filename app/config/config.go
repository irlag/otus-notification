package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	BindAddress string `envconfig:"optional"`
	Port        string `envconfig:"PORT"`
	Debug       bool   `envconfig:"DEBUG"`
}

func NewConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	var err error

	if env != "" {
		err = godotenv.Load(".env." + env + ".local")
		if err != nil {
			return nil, err
		}
	}

	if "test" != env {
		err = godotenv.Load(".env.local")
		if err != nil {
			return nil, err
		}
	}
	if env != "" {
		err = godotenv.Load(".env." + env)
		if err != nil {
			return nil, err
		}
	}

	err = godotenv.Load() // The Original .env
	if err != nil {
		return nil, err
	}

	config := &Config{
		Port:  "8000",
		Debug: false,
	}

	if err := envconfig.Init(config); err != nil {
		return nil, err
	}

	return config, nil
}
