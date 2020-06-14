package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func Initialize() {
	viper.SetConfigName("config\\config")
	viper.SetConfigType("yaml")
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Error reading config.yaml: %s", err))
	}
}

func GetBaseUrl() string {
	return viper.GetString("BaseUrl")
}

func GetClient() string {
	return viper.GetString("ClientId")
}

func GetSecret() string {
	return viper.GetString("ClientSecret")
}

func GetDBConnection() string {
	return viper.GetString("DbConnection")
}
