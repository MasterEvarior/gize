package helper

import (
	"log"
	"os"
	"strings"
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

func IsEnabled(name string, defaultValue bool) bool {
	value, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}

	if strings.ToLower(value) == "true" || value == "1" || strings.ToLower(value) == "yes" {
		return !defaultValue
	}

	return defaultValue
}
