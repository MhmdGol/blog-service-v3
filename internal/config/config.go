package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	HttpURI       string             `mapstructure:"HTTP_URI"`
	Port          string             `mapstructure:"PORT"`
	SecretKey     string             `mapstructure:"SECRET_KEY"`
	SQLDatabase   SQLDatabaseConfig  `mapstructure:",squash"`
	NoSQLDatabase NoSQLDtabaseConfig `mapstructure:",squash"`
}

type SQLDatabaseConfig struct {
	Host     string `mapstructure:"SQL_DATABASE_HOST"`
	Port     string `mapstructure:"SQL_DATABASE_PORT"`
	User     string `mapstructure:"SQL_DATABASE_USER"`
	Password string `mapstructure:"SQL_DATABASE_PASSWORD"`
	Name     string `mapstructure:"SQL_DATABASE_NAME"`
	Ssl      string `mapstructure:"SQL_DATABASE_SSL"`
}

type NoSQLDtabaseConfig struct {
	Host string `mapstructure:"NOSQL_DATABASE_HOST"`
	Port string `mapstructure:"NOSQL_DATABASE_PORT"`
	Name string `mapstructure:"NOSQL_DATABASE_NAME"`
}

func Load() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	viper.BindEnv(
		"HTTP_URI",
		"PORT",
		"SECRET_KEY",
		"SQL_DATABASE_HOST",
		"SQL_DATABASE_PORT",
		"SQL_DATABASE_USER",
		"SQL_DATABASE_PASSWORD",
		"SQL_DATABASE_NAME",
		"SQL_DATABASE_SSL",
		"NOSQL_DATABASE_HOST",
		"NOSQL_DATABASE_PORT",
	)

	var c Config
	err = viper.Unmarshal(&c)

	return c, err
}
