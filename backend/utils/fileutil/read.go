package fileutil

import (
	"fmt"
	"io"
	"os"

	"implude.kr/VOAH-Official-File/configs"
	awsutil "implude.kr/VOAH-Official-File/utils/awsuitl"
)

func ReadFileFromLocal(fileID int) (io.Reader, error) {
	dataDir := configs.Env.Server.DataDir
	f, err := os.Open(fmt.Sprintf("%s/%d.blob", dataDir, fileID))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func ReadFileFromS3(fileID int) ([]byte, error) {
	return awsutil.DownloadFileFromS3(fmt.Sprintf("%d", fileID))
}
