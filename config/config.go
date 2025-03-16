package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Port               int    `env:"PORT" envdefault:"8080"`
	Enviroment         string `env:"ENV" envdefault:"development"`
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID,required"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET,required"`
	RandomState        string `env:"RANDOM_STATE" envdefault:"random_state"`
	HostURL            string `env:"DOMAIN_HOST_URL" envdefault:"localhost:8080"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load environment")
	}

	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
