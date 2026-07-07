package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server struct {
		Listen string `json:"listen"`
	} `json:"server"`

	Kea struct {
		URL string `json:"url"`
	} `json:"kea"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)

	err = json.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
