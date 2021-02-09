package setuputil

import (
	"github.com/spf13/viper"
)

//SetupConfig setup viper to return config options
func SetupConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	err := config.ReadInConfig()

	if err != nil {
		panic("Viper setup failed")
	}

	return config
}
