package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env    string `mapstructure:"env"`
	Port   int    `mapstructure:"app_port"`
	DBSlot int    `mapstructure:"db_slot"`
	DBUser string `mapstructure:"db_user"`
	DBPass string `mapstructure:"db_pass"`
	DBAddr string `mapstructure:"db_addr"`
}

func LoadConfig() Config {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetDefault("env", "production")
	viper.SetDefault("db_slot", 0)
	viper.SetDefault("db_addr", "localhost:6379")
	viper.SetDefault("app_port", 1337)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, using env variables")
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("Unable to decode config:", err)
	}

	return cfg
}
