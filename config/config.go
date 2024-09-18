package config

import (
	"encoding/json"
	"os"
)

const ConfigPath = "../config.json"

type Config struct {
	Api Api    `json:"api"`
	Db  string `json:"db"`
}

type Api struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

func GetConfig() (*Config, error) {
	configBytes, err := os.ReadFile(ConfigPath)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
