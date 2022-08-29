package LibUtils

import (
	"fmt"
	"github.com/spf13/viper"
  )

 func InitConfi() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// viper.SetDefault("database.dbname", "test_db")
	// var configuration c.Configurations
	// err := viper.Unmarshal(&configuration)
	// if err != nil {
	// 	fmt.Printf("Unable to decode into struct, %v", err)
	// }
	
 }