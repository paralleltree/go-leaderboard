package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	Port          int    `default:"8000"`
	RedisEndpoint string `split_words:"true" required:"true"`
	AllowOrigin   string `split_words:"true"`
}

func GetEnv() Env {
	env := Env{}
	if err := envconfig.Process("", &env); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}
	return env
}
