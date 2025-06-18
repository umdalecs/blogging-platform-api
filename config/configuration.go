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
		DbUser:     utils.LoadEnv("MYSQL_DB_USER", "root"),
		DbPassw:    utils.LoadEnv("MYSQL_DB_PASSWD", "root"),
		DbName:     utils.LoadEnv("MYSQL_DB_NAME", "bloggingPlatform"),
		DbAddr:     utils.LoadEnv("MYSQL_DB_ADDR", "127.0.0.1:3306"),
	}
}
