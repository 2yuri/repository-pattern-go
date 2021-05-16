package config

import (
	"log"

	"github.com/spf13/viper"
)

type App struct {
	ServerHost              string `mapstructure:"host"`
	ServerPort              string `mapstructure:"port"`
	DatabaseHost            string `mapstructure:"db_host"`
	DatabasePort            string `mapstructure:"db_port"`
	DatabaseSslMode         string `mapstructure:"db_ssl_mode"`
	DatabaseUser            string `mapstructure:"db_user"`
	DatabaseName            string `mapstructure:"db_name"`
	DatabasePass            string `mapstructure:"db_pass"`
	DatabaseType            string `mapstructure:"db_type"`
	DatabaseMaxIdleConns    int    `mapstructure:"db_max_idle_conns"`
	DatabaseMaxOpensConns   int    `mapstructure:"db_max_opens_conns"`
	DatabaseConnMaxLifetime int64  `mapstructure:"db_conn_max_lifetime"`
}

var envConfig *viper.Viper
var config *App

func Init() {

	envConfig = viper.New()
	envConfig.AddConfigPath(".")
	envConfig.AddConfigPath("../")
	envConfig.SetConfigType("env")
	envConfig.SetConfigName(`.env`)
	if err := envConfig.ReadInConfig(); err != nil {
		log.Fatalf("Error on reading the envConfig file: %v", err)
	}
	marshallErr := envConfig.Unmarshal(&config)
	if marshallErr != nil {
		log.Fatalf("Error on unmarshalling the envConfig file: %v", marshallErr)
	}

}

func GetConfig() *App {
	return config
}
