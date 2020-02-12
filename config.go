package main

import "github.com/spf13/viper"

func SetupConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")

	// default values
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", 5432)
	viper.SetDefault("db.username", "")
	viper.SetDefault("db.password", "")
	viper.SetDefault("db.database", "baconpants")

	return viper.ReadInConfig()
}
