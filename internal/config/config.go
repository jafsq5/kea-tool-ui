package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Listen string `yaml:"listen"`
	} `yaml:"server"`

	Kea struct {
		URL string `yaml:"url"`
	} `yaml:"kea"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	err = yaml.Unmarshal(data, cfg)

	return cfg, err
}
