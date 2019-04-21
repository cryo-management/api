package config

import (
	"github.com/BurntSushi/toml"
)

// Config defines the struct of system configuration parameters
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// Load get the system configuration parameters from the file and put on struct
func Load() (Config, error) {
	file := "config.toml"
	config := Config{}
	_, err := toml.DecodeFile(file, &config)

	return config, err
}
