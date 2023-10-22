package file

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/models"
	"implude.kr/VOAH-Official-File/utils/fileutil"
	"implude.kr/VOAH-Official-File/utils/validator"
)

type UploadFileRequest struct {
	FileName          string `validate:"required,min=1,max=100"`
	FileType          string `validate:"required,min=1,max=100"`
	AffiliationType   string `validate:"required,min=1,max=100"`
	AffiliationTarget string `validate:"required,uuid4"`
}

func UploadFileCtrl(c *fiber.Ctx) error {
	// read form data
	uploadRequest := UploadFileRequest{
		FileName:          c.FormValue("filename", ""),
		FileType:          c.FormValue("filetype", ""),
		AffiliationType:   c.FormValue("affiliation-type", ""),
		AffiliationTarget: c.FormValue("affiliation-target", ""),
	}
	if errArr := validator.VOAHValidator.Validate(uploadRequest); len(errArr) > 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   "File is required",
		})
	}

	newFile := &models.File{
		OwnerID:   uuid.New(),
		FileName:  uploadRequest.FileName,
		FileType:  uploadRequest.FileType,
		Processed: 1, // 1: Processing, 2: Processed, 3: Failed
	}
	db := database.DB
	db.Create(&newFile)

	fileStoreConf := configs.Env.File
	fileName := fmt.Sprintf("%d.blob", newFile.ID)
	filePath := fmt.Sprintf("%s/%s", fileStoreConf.TempDataDir, fileName)
	c.SaveFile(fileHeader, filePath)

	go fileutil.SaveFilePermanently(filePath, fileName, newFile)

	return c.JSON(fiber.Map{
		"message": "File is on Processing",
		"file-id": newFile.ID,
	})
}
