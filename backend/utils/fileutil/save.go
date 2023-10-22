package fileutil

import (
	"fmt"
	"io"
	"os"

	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/models"
	awsutil "implude.kr/VOAH-Official-File/utils/awsuitl"
	"implude.kr/VOAH-Official-File/utils/logger"
)

func SaveFileToS3(tempFilePath string, fileName string, metadata *models.File) {
	log := logger.Logger
	f, err := os.Open(tempFilePath)
	if err != nil {
		log.Error(err.Error())
		return
	}

	var fileReader io.Reader = f
	err = awsutil.UploadFileToS3(fileReader, fmt.Sprintf("%d", metadata.ID))
	f.Close()
	os.Remove(tempFilePath)
	if err != nil {
		log.Error(err.Error())
		metadata.Processed = 3
	} else {
		metadata.Processed = 2
	}
	db := database.DB
	err = db.Save(&metadata).Error
	if err != nil {
		log.Error(err.Error())
	}
}
