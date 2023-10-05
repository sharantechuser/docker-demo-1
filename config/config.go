package config

import (
	"os"
)

var User = getEnv("USER_NAME", "")

func getEnv(key, fallback string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		value = fallback
	}
	return value
}
