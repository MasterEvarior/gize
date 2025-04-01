package helper

import (
	"log"
	"os"
)

func GetEnvVar(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("Environment variable '%s' was not defined", name)
	}
	return value
}

func GetEnvVarWithDefault(name string, defaultValue string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	return value
}
