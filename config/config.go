package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const ConfigPath = "./config.json"

type Config struct {
	Api Api    `json:"api"`
	Db  string `json:"db"`
}

type Api struct {
	Host string `json:"host"`
	Port uint16 `json:"port"`
}

func GetProjectDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}

	return filepath.Dir(exePath), nil
}

func GetConfig() (*Config, error) {
	projectDir, err := GetProjectDir()
	if err != nil {
		return nil, err
	}

	configBytes, err := os.ReadFile(filepath.Join(projectDir, ConfigPath))
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
