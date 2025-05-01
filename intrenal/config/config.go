package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env    string `env:"ENV"`
	ApiKey string `env:"API_KEY"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		fmt.Errorf("config file not found: %w, using environment variables", err)
	}

	return &cfg, nil
}
