package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Host  string `mapstructure:"host"`
	Token string `mapstructure:"token"`
}

var configDir = ""

func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir = filepath.Join(home, ".config", "gt")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	viper.SetEnvPrefix("GT")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	return nil
}

func Get() (*Config, error) {
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func Set(key, value string) error {
	viper.Set(key, value)
	return viper.WriteConfig()
}

func GetHost() string {
	return viper.GetString("host")
}

func GetToken() string {
	return viper.GetString("token")
}
