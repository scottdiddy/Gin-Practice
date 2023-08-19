//Package config helps to set the simple configuration for using viper to read files
//that might be in a different format like json, yaml etc.
package config

import (
	"log"
	"github.com/spf13/viper"
)

var Appconfig *viper.Viper

func Init()  {
	Appconfig = viper.New()
	Appconfig.SetConfigType("yaml")
	Appconfig.SetConfigName("config")
	Appconfig.AddConfigPath("src/config/")
	err := Appconfig.ReadInConfig()
	if err != nil {
		log.Fatal("error on passing configuration file", err.Error())
	}
}