package config

import "os"

// all the functions in this config file can be modified to use viper
func GetJsonApiHost() string {
	return os.Getenv("JSON_API_HOST")
}
