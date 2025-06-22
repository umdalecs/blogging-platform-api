package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
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
		DbUser:  LoadEnv("MYSQL_DB_USER"),
		DbPassw: LoadEnv("MYSQL_DB_PASSWD"),
		DbName:  LoadEnv("MYSQL_DB_NAME"),
		DbAddr:  LoadEnv("MYSQL_DB_ADDR"),
	}
}

func LoadEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("missing %s", name)
	}

	return value
}

func LoadIntEnv(name string) int {
	value, err := strconv.Atoi(LoadEnv(name))
	if err != nil {
		log.Fatalf("%s must be an integer", name)
	}

	return value
}
