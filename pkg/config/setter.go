package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("error reading the configs")
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatal("unable to decode into struct")
	}

}
