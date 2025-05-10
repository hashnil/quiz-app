package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Load loads the application configuration.
func Load() error {
	// Load configuration from local YAML file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./infrastructure/config/")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read local config: %v", err)
	}

	// Enable automatic environment variable binding
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}
