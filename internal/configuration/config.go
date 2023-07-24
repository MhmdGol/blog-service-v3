package configuration

import (
	"fmt"

	"github.com/spf13/viper"
)

func SetConfigs() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// path set from main.go reference
	viper.AddConfigPath("./../internal/configuration")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
}
