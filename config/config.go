package config

import (
	"log"

	"github.com/spf13/viper"
)

var AppConfig Config

type Config struct {
	PORT                       int
	MONGO_HOST                 string
	MONGO_PORT                 string
	MONGO_USERNAME             string
	MONGGO_PASSWORD            string
	MONGO_DATABASE             string
	MONGO_POOL_MIN             int
	MONGO_POOL_MAX             int
	MONGO_MAX_IDLE_TIME_SECOND int
}

func InitializeAppConfig() error {
	viper.SetConfigName(".env") // allow directly reading from .env file
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("/")
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	AppConfig.PORT = viper.GetInt("PORT")

	AppConfig.MONGO_HOST = viper.GetString("MONGO_HOST")
	AppConfig.MONGO_PORT = viper.GetString("MONGO_PORT")
	AppConfig.MONGO_USERNAME = viper.GetString("MONGO_USERNAME")
	AppConfig.MONGGO_PASSWORD = viper.GetString("MONGO_PASSWORD")
	AppConfig.MONGO_DATABASE = viper.GetString("MONGO_DATABASE")
	AppConfig.MONGO_POOL_MIN = viper.GetInt("MONGO_POOL_MIN")
	AppConfig.MONGO_POOL_MAX = viper.GetInt("MONGO_POOL_MAX")
	AppConfig.MONGO_MAX_IDLE_TIME_SECOND = viper.GetInt("MONGO_MAX_IDLE_TIME_SECOND")

	log.Println("[INIT] configuration loaded")
	return nil
}
