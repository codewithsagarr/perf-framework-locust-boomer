package utils

import "os"

// Used to get the value of the given environment variable. If the environment variable is not found retrun the defaultValue.
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
