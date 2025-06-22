package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Envs = initConfig()

type Configuration struct {
	DbUser string
	DbPass string
	DbName string
	DbAddr string
	DbPort uint16
}

func initConfig() *Configuration {
	godotenv.Load()

	return &Configuration{
		DbUser: loadEnv("PG_DB_USER"),
		DbPass: loadEnv("PG_DB_PASS"),
		DbName: loadEnv("PG_DB_NAME"),
		DbAddr: loadEnv("PG_DB_ADDR"),
		DbPort: uint16(loadIntEnv("PG_DB_PORT")),
	}
}

func loadEnv(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("missing %s", name)
	}

	return value
}

func loadIntEnv(name string) int {
	value, err := strconv.Atoi(loadEnv(name))
	if err != nil {
		log.Fatalf("%s must be an integer", name)
	}

	return value
}
