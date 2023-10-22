package fileutil

import (
	"fmt"
	"os"

	"implude.kr/VOAH-Official-File/configs"
	awsutil "implude.kr/VOAH-Official-File/utils/awsuitl"
)

func DeleteFileFromS3(fileID int) error {
	awsutil.DeleteFileFromS3(fmt.Sprintf("%d", fileID))
	return nil
}

func DeleteFileFromLocal(fileID int) error {
	dataDir := configs.Env.Server.DataDir
	return os.Remove(fmt.Sprintf("%s/%d.blob", dataDir, fileID))
}
