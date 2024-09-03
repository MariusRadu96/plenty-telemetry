package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v1"
)

type Config struct {
	LogLevel string         `yaml:"log_level"`
	Drivers  []DriverConfig `yaml:"drivers"`
}

type DriverConfig struct {
	Type       string            `yaml:"type"`
	Attributes map[string]string `yaml:"attributes,omitempty"`
}

func LoadConfg(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var config Config
	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %v", err)
	}

	return &config, nil
}
