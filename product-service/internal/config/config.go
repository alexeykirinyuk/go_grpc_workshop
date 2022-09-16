package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"time"
)

type DB struct {
	DSN             string        `yaml:"DSN"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxIdleTime time.Duration `yaml:"connMaxIdleTime"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

func (db DB) GetDSN() string {
	return db.DSN
}

func (db DB) GetMaxOpenConns() int {
	return db.MaxOpenConns
}

func (db DB) GetMaxIdleConns() int {
	return db.MaxIdleConns
}

func (db DB) GetConnMaxIdleTime() time.Duration {
	return db.ConnMaxIdleTime
}

func (db DB) GetConnMaxLifetime() time.Duration {
	return db.ConnMaxLifetime
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
