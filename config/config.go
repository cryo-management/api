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

// NewConfig get the system configuration parameters from the file and put on struct
func NewConfig(file string) (Config, error) {
	config := Config{}
	_, err := toml.DecodeFile(file, &config)

	return config, err
}
