package config

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Server ServerConfig `json:"server"`
	Kea    KeaConfig    `json:"kea"`
	SSH    SSHConfig    `json:"ssh"`
}

type ServerConfig struct {
	Listen string `json:"listen"`
}

type KeaConfig struct {
	ControlAgent string `json:"control-agent"`
	HostsFile    string `json:"hosts-file"`
}

type SSHConfig struct {
	Host       string        `json:"host"`
	Port       int           `json:"port"`
	User       string        `json:"user"`
	PrivateKey string        `json:"private-key"`
	Timeout    time.Duration `json:"timeout"`
}

func Load(path string) (*Config, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		SSH: SSHConfig{
			Port:    22,
			Timeout: 30 * time.Second,
		},
	}

	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}

	if cfg.SSH.Port == 0 {
		cfg.SSH.Port = 22
	}

	if cfg.SSH.Timeout == 0 {
		cfg.SSH.Timeout = 30 * time.Second
	}

	return cfg, nil
}
