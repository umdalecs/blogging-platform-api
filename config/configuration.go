package config

import (
	"github.com/joho/godotenv"
	"github.com/umdalecs/blogging-platform-api/utils"
)

var Envs = initConfig()

type Configuration struct {
	DbUser  string
	DbPassw string
	DbName  string
	DbAddr  string
}

func initConfig() *Configuration {
	godotenv.Load()

	return &Configuration{
		DbUser:  utils.LoadEnv("MYSQL_DB_USER"),
		DbPassw: utils.LoadEnv("MYSQL_DB_PASSWD"),
		DbName:  utils.LoadEnv("MYSQL_DB_NAME"),
		DbAddr:  utils.LoadEnv("MYSQL_DB_ADDR"),
	}
}
