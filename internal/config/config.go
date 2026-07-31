package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

/*
MustLoad() env resolution:

  dev.env exists?
    NO  -> APP_ENV set?
             NO  -> error: dev.env not found and APP_ENV is not set
             YES -> file = "<APP_ENV>.env"
    YES -> file = "dev.env"

  file exists?
    NO  -> error: "<file> not found"
    YES -> load file into cfg

  loaded via dev.env AND cfg.Env != "dev"?
    YES -> file = "<cfg.Env>.env"
             exists? NO -> error: "<file> not found"
             exists? YES -> reload file into cfg
    NO  -> skip

  validate cfg (Env in {dev,test,prod}, Port required) -> return cfg or error
*/

type Config struct {
	Env  string `env:"ENV" env-default:"dev" validate:"required,oneof=dev test prod"`
	Port string `env:"PORT" env-default:"8000" validate:"required"`
}

const bootstrapFile = "dev.env"

func MustLoad() (*Config, error) {

	cfg := &Config{}

	file := bootstrapFile

	if _, err := os.Stat(bootstrapFile); err != nil {
		// no dev.env on disk (e.g. prod host) — fall back to APP_ENV
		env := os.Getenv("APP_ENV")
		if env == "" {
			return nil, fmt.Errorf("config file %s not found and APP_ENV is not set", bootstrapFile)
		}
		file = fmt.Sprintf("%s.env", env)
	}

	if _, err := os.Stat(file); err != nil {
		return nil, fmt.Errorf("config file %s not found: %w", file, err)
	}

	if err := cleanenv.ReadConfig(file, cfg); err != nil {
		return nil, err
	}

	// dev.env's ENV key decides which env is actually active;
	// if it points elsewhere (e.g. prod), reload from that file.
	if file == bootstrapFile && cfg.Env != "dev" {
		file = fmt.Sprintf("%s.env", cfg.Env)

		if _, err := os.Stat(file); err != nil {
			return nil, fmt.Errorf("config file %s not found: %w", file, err)
		}

		if err := cleanenv.ReadConfig(file, cfg); err != nil {
			return nil, err
		}
	}

	//validate
	validator := validator.New()
	if err := validator.Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil

}
