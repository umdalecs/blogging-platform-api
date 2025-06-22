package utils

import (
	"log"
	"os"
	"strconv"
)

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
