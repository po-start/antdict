package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	APP_HASH_KEY	string	`yaml:"APP_HASH_KEY"`
	APP_NAME		string	`yaml:"APP_NAME"`
	APP_DEBUG		bool	`yaml:"APP_DEBUG"`
	APP_TIMEZONE	string	`yaml:"APP_TIMEZONE"`
	APP_ENV			string	`yaml:"APP_ENV"`

	USER			string	`yaml:"USER"`
	PASSWORD		string	`yaml:"PASSWORD"`

	ADDRESS			string	`yaml:"ADDRESS"`
	PROT			int32	`yaml:"PROT"`

	LOG_DEBUG		bool	`yaml:"LOG_DEBUG"`
	LOG_PATH		string	`yaml:"LOG_PATH"`
	LOG_LEVEL		string	`yaml:"LOG_LEVEL"`
	LOG_FILE		string	`yaml:"LOG_FILE"`
	LOG_SIZE		int		`yaml:"LOG_SIZE"`

	DB_CONNECTION	string	`yaml:"DB_CONNECTION"`
	DB_HOST			string	`yaml:"DB_HOST"`
	DB_PORT			int32	`yaml:"DB_PORT"`
	DB_DATABASE		string	`yaml:"DB_DATABASE"`
	DB_USERNAME		string	`yaml:"DB_USERNAME"`
	DB_PASSWORD		string	`yaml:"DB_PASSWORD"`
	DB_CHARSET		string	`yaml:"DB_CHARSET"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return ParseConfigData(data)
}

func WriteConfigFile(cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	execPath, err := os.Getwd()
	if err != nil {
		return err
	}

	configPath := execPath + "etc/config.yaml"
	err = ioutil.WriteFile(configPath, data, 0755)
	if err != nil {
		return err
	}
	return nil
}
