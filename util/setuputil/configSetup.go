package setuputil

import (
	viper "github.com/spf13/viper"
)

//SetupConfig setup viper to return config options
func SetupConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")
	config.AddConfigPath("E:\\Projects\\GoProjects\\simpleWebGolang")
	err := config.ReadInConfig()

	if err != nil {
		panic("Viper setup failed: " + err.Error())
	}

	return config
}
