package main

import (
	"github.com/caarlos0/env/v6"
)

type Environment string

const (
	// EnvironmentLocal is the local environment.
	EnvironmentLocal Environment = "local"
	// EnvironmentAWS is the AWS environment.
	EnvironmentAWS Environment = "aws"
)

type Config struct {
	// Environment is the environment the service is running in.
	Environment Environment `env:"ENVIRONMENT" envDefault:"local"`
}

// NewConfig returns a new configuration object with environment variables loaded.
func NewConfig() (Config, error) {
	cfg := Config{}

	return cfg, env.Parse(&cfg)
}
