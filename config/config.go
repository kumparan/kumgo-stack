package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// HostName :nodoc:
func HostName() string {
	return viper.GetString("hostname")
}

// Port :nodoc:
func Port() string {
	return viper.GetString("port")
}

// Env :nodoc:
func Env() string {
	return viper.GetString("env")
}

// HostName :nodoc:
func AppName() string {
	return viper.GetString("app_name")
}

// PapertrailEndpoint :nodoc:
func PapertrailEndpoint() string {
	return viper.GetString("papertrail_endpoint")
}

// ArangoHost :nodoc:
func ArangoHost() string {
	return viper.GetString("arangodb.host")
}

// ArangoDatabase :nodoc:
func ArangoDatabase() string {
	return viper.GetString("arangodb.database")
}

// ArangoUsername :nodoc:
func ArangoUsername() string {
	return viper.GetString("arangodb.username")
}

// ArangoPassword :nodoc:
func ArangoPassword() string {
	return viper.GetString("arangodb.password")
}

// CollectionStory :nodoc:
func CollectionStory() string {
	return viper.GetString("arangodb.collections.stories")
}

// GetConf :nodoc:
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
