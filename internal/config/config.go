package config

import (
	"encoding/json"
	"os"
)

type Config struct {
    Server ServerConfig `json:"server"`
    Kea    KeaConfig    `json:"kea"`
}

type KeaConfig struct {
    ControlAgent string `json:"control-agent"`
    HostsFile    string `json:"hosts-file"`
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
