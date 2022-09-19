package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type DB struct {
	DSN string `yaml:"dsn"`
}

type Config struct {
	DB DB `yaml:"db"`
}

var cfg *Config

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	if cfg != nil {
		return nil
	}

	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return err
	}

	return nil
}

func GetConfigInstance() Config {
	if cfg != nil {
		return *cfg
	}

	return Config{}
}
