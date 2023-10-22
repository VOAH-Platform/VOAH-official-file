package configs

import (
	"fmt"
	"os"
	"strconv"
)

func getEnvStr(key string, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		fmt.Printf("Environment variable %s is not set, Keep Going with Default Value '%s' \n", key, defaultValue)
		return defaultValue
	}
	return
}

func getEnvInt(key string, defaultValue int) (intValue int) {
	value := getEnvStr(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic(err)
	}
	return
}

func LoadEnv() {
	Env = &MainEnv{
		Server: serverEnv{
			HostURL:    getEnvStr("SERVER_HOST_URL", "http://localhost:3002"),
			Port:       getEnvInt("SERVER_PORT", 3002),
			CSRFOrigin: getEnvStr("SERVER_CSRF_ORIGIN", "*"),
			DataDir:    getEnvStr("SERVER_DATA_DIR", "./data"),
		},
		File: fileStoreEnv{
			USE_S3:          getEnvStr("FILE_USE_S3", "false") == "true",
			FileSizeLimitMB: getEnvInt("FILE_SIZE_LIMIT_MB", 1024),
			TempDataDir:     getEnvStr("FILE_TEMP_DATA_DIR", "./temp_data"),
		},
		Database: databaseEnv{
			Host:     getEnvStr("DB_HOST", "localhost"),
			Port:     getEnvInt("DB_PORT", 5432),
			User:     getEnvStr("DB_USER", "postgres"),
			Password: getEnvStr("DB_PASSWORD", "password"),
			DBName:   getEnvStr("DB_NAME", "voah-file"),
		},
	}
	if Env.File.USE_S3 {
		Env.File.S3 = s3Env{
			AcessKeyID:      getEnvStr("FILE_S3_ACCESS_KEY_ID", ""),
			SecretAccessKey: getEnvStr("FILE_S3_SECRET_ACCESS_KEY", ""),
			Bucket:          getEnvStr("FILE_S3_BUCKET", ""),
			Region:          getEnvStr("FILE_S3_REGION", ""),
		}
	}
}
