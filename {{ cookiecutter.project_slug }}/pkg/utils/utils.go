package utils

import (
	"os"
	"strconv"
)

func LookupEnvOrString(key string, defaultValue string) string {
	envVariable, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return envVariable
}

func LookupEnvOrInt(key string, defaultValue int) int {
	envVariable, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	num, err := strconv.Atoi(envVariable)
	if err != nil {
		panic(err.Error())
	}
	return num
}
