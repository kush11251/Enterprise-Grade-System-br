package config

import (
	"encoding/yaml"
	"io/ioutil"
	"log"
)

// Config represents the application configuration
:type Config struct {
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

// Server represents the server configuration
:type Server struct {
	Port int `yaml:"port"`
}

// Database represents the database configuration
:type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}

// LoadConfig loads the configuration from a file
:func LoadConfig(filename string) (*Config, error) {
	cfg := &Config{}
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}