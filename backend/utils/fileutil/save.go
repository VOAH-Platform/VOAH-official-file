package fileutil

import (
	"fmt"
	"io"
	"os"

	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/models"
	awsutil "implude.kr/VOAH-Official-File/utils/awsuitl"
	"implude.kr/VOAH-Official-File/utils/logger"
)

func SaveFilePermanently(tempFilePath string, fileName string, metadata *models.File) {
	fileConf := configs.Env.File

	var err error
	if fileConf.USE_S3 {
		err = saveFileToS3(tempFilePath, fileName, fmt.Sprintf("%d", metadata.ID))
	} else {
		err = saveFileToLocal(tempFilePath, fileName)
	}
	db := database.DB
	log := logger.Logger
	if err != nil {
		log.Error(err.Error())
		metadata.Processed = 3
		db.Save(&metadata)
		return
	}
	metadata.Processed = 2
	err = db.Save(&metadata).Error
	if err != nil {
		log.Error(err.Error())
	}
}

func saveFileToLocal(tempFilePath string, fileName string) error {
	dataDir := configs.Env.Server.DataDir
	// move file to data dir
	return os.Rename(tempFilePath, dataDir+"/"+fileName)
}

func saveFileToS3(tempFilePath string, fileName string, fileID string) error {
	// s3Conf := configs.Env.File.S3

	f, err := os.Open(tempFilePath)
	if err != nil {
		return err
	}

	var fileReader io.Reader = f
	err = awsutil.UploadFileToS3(fileReader, fileID)
	f.Close()
	os.Remove(tempFilePath)

	return err
}
