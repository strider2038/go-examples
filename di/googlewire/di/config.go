package di

import "github.com/caarlos0/env/v11"

type Config struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func NewConfig() (*Config, error) {
	config, err := env.ParseAs[Config]()

	return &config, err
}
