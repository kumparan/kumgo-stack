package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func HostName() string {
	return viper.GetString("hostname")
}

func Port() string {
	return viper.GetString("port")
}

func Env() string {
	return viper.GetString("env")
}

func AppName() string {
	return viper.GetString("app_name")
}

func PapertrailEndpoint() string {
	return viper.GetString("papertrail_endpoint")
}

func ArangoHost() string {
	return viper.GetString("arangodb.host")
}

func CollectionStory() string {
	return viper.GetString("arangodb.collections.stories")
}

func GetConf() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetEnvPrefix("svc")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("%v", err)
	}

	return
}
