package directory

import (
	"os"

	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/utils/logger"
)

func IniteDirectory() {
	createDirIfNotExist(configs.Env.File.TempDataDir)
	if !configs.Env.File.USE_S3 {
		createDirIfNotExist(configs.Env.Server.DataDir)
	}
}

func createDirIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log := logger.Logger
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Directory created: " + path)
	}
}
