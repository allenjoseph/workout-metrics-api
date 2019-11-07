package config

import (
	"os"

	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/spf13/viper"
)

// LoadConfig function
func LoadConfig() {
	workingDir, err := os.Getwd()
	util.FailOnError("Failed to load config", err)

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workingDir)

	err = viper.ReadInConfig()
	util.FailOnError("Failed to load config", err)
}
