package configs

import (
	"os"
	"strconv"
)

func getEnvStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	value := getEnvStr(key, "")
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}

func LoadEnv() {
	Env = &MainEnv{
		Server: serverEnv{
			HostURL:    getEnvStr("SERVER_HOST_URL", "http://localhost:3000"),
			Port:       getEnvInt("SERVER_PORT", 3000),
			CSRFOrigin: getEnvStr("SERVER_CSRF_ORIGIN", "*"),
			DataDir:    getEnvStr("SERVER_DATA_DIR", "./data"),
		},
	}
}
