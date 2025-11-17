package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	jsonPath, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}
	file, err := os.Open(jsonPath)
	if err != nil {
		return Config{}
	}
	defer file.Close()
	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}
	}

	return cfg

}

func (c *Config) SetUser(username string) {
	c.CurrentUserName = username
	write(*c)
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := home + "/" + configFileName
	return path, nil

}

func write(cfg Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}
	os.WriteFile(path, jsonData, 0644)
	return nil
}
