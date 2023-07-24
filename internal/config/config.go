package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	HttpURI  string         `mapstructure:"HTTP_URI"`
	Database DatabaseConfig `mapstructure:",squash"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     string `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABASE_USER"`
	Password string `mapstructure:"DATABASE_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
	Ssl      string `mapstructure:"DATABASE_SSL"`
}

func Load() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	// path set from main.go reference
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	viper.BindEnv("HTTP_URI",
		"DATABASE_HOST",
		"DATABASE_PORT",
		"DATABASE_USER",
		"DATABASE_PASSWORD",
		"DATABASE_NAME",
		"DATABASE_SSL",
	)

	var c Config
	err = viper.Unmarshal(&c)

	return c, err
}
