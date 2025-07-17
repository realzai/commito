package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Provider string `json:"provider"`
	Model    string `json:"model"`
	ApiKey   string `json:"api_key"`
}

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".commito", "config.json"), nil
}

func SaveConfig(cfg Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(configPath), 0755)
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(cfg)
}

func LoadConfig() (Config, error) {
	var cfg Config
	configPath, err := getConfigPath()
	if err != nil {
		return cfg, err
	}
	file, err := os.Open(configPath)
	if err != nil {
		return cfg, err
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&cfg)
	return cfg, err
}
