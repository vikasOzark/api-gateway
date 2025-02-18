package helpers

import (
	"os"
	"strconv"
)

func ConvertEnvInt(key string) (int, error) {
	val := os.Getenv(key)
	value, err := strconv.Atoi(val)
	return value, err
}

func ConvertInt(val string) (int, error) {
	value, err := strconv.Atoi(val)
	return value, err
}
