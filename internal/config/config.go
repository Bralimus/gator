package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	JSONFile, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(JSONFile)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	JSONFile := home + "/" + configFileName

	return JSONFile, nil

}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	return write(*cfg)
}

func write(cfg Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
