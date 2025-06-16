package utils

import (
	"log"
	"os"
	"strconv"
)

func LoadEnv(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}

	return fallback
}

func LoadIntEnv(name, fallback string) int {
	val, err := strconv.Atoi(LoadEnv(name, fallback))
	if err != nil {
		log.Fatal("Missconfig envs")
	}

	return val
}
