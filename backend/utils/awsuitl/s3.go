package awsutil

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/utils/logger"
)

func newS3Client() (*s3.Client, error) {
	log := logger.Logger
	credntial := credentials.NewStaticCredentialsProvider(configs.Env.File.S3.AcessKeyID, configs.Env.File.S3.SecretAccessKey, "")
	cfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithCredentialsProvider(credntial),
		awsConfig.WithRegion(configs.Env.File.S3.Region),
	)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return s3.NewFromConfig(cfg), nil
}

func UploadFileToS3(uploadFile io.Reader, key string) error {
	log := logger.Logger
	client, err := newS3Client()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	uploader := manager.NewUploader(client)
	_, err = uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(configs.Env.File.S3.Bucket),
		Key:    aws.String(key),
		Body:   uploadFile,
	})
	if err != nil {
		log := logger.Logger
		log.Error(err.Error())
		return err
	}
	return nil
}

func DownloadFileFromS3(key string) ([]byte, error) {
	buffer := manager.NewWriteAtBuffer([]byte{})

	log := logger.Logger
	client, err := newS3Client()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	downloader := manager.NewDownloader(client)
	_, err = downloader.Download(context.TODO(), buffer, &s3.GetObjectInput{
		Bucket: aws.String(configs.Env.File.S3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return buffer.Bytes(), nil
}

func DeleteFileFromS3(key string) error {
	log := logger.Logger
	client, err := newS3Client()
	if err != nil {
		log.Error(err.Error())
		return err
	}
	_, err = client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(configs.Env.File.S3.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
