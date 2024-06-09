package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type ServerConfig struct {
	Port               int           `yaml:"port"`
	ReadTimeout        time.Duration `yaml:"read_timeout"`
	WriteTimeout       time.Duration `yaml:"write_timeout"`
	IdleTimeout        time.Duration `yaml:"idle_timeout"`
}

type DatabaseConfig struct {
	FilePath          string        `yaml:"file_path"`
	ConnectionTimeout time.Duration `yaml:"connection_timeout"`
	ConnMaxLifetime   time.Duration `yaml:"conn_max_lifetime"`
	MaxOpenConns      int           `yaml:"max_open_conns"`
	MaxIdleConns      int           `yaml:"max_idle_conns"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig() *Config {
	file, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}
	var config Config
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatalf("Unable to decode config into struct: %v", err)
	}

	return &config
}
