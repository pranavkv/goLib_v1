/*
Author: Pranav KV
Mail: pranavkvnambiar@gmail.com
*/
package golib_v1

import (
	"fmt"

	config "github.com/spf13/viper"
)

func InitConfig(path string) {
	config.SetConfigName("config")
	config.AddConfigPath(".")
	config.AutomaticEnv()

	config.SetConfigType("yml")
	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		config.SetConfigType("properties")
		_ = config.ReadInConfig()
	}

}

//TODO: throw key not found exception

func GetString(key string) string {
	return config.GetString(key)
}

func GetInt(key string) int {
	return config.GetInt(key)
}

func GetBoolean(key string) bool {
	return config.GetBool(key)
}
