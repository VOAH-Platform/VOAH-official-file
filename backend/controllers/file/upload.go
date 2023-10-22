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
	err = db.Create(&newFile).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	fileName := fmt.Sprintf("%d.blob", newFile.ID)
	var filePath string
	if configs.Env.File.USE_S3 {
		filePath = fmt.Sprintf("%s/%s", configs.Env.File.TempDataDir, fileName)
	} else {
		filePath = fmt.Sprintf("%s/%s", configs.Env.Server.DataDir, fileName)
	}

	err = c.SaveFile(fileHeader, filePath)
	if err != nil {
		newFile.Processed = 3
		db.Save(&newFile)
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	if configs.Env.File.USE_S3 {
		go fileutil.SaveFileToS3(filePath, fileName, newFile)
	} else {
		newFile.Processed = 2
		if err = db.Save(&newFile).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal Server Error",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "File is on Processing",
		"file-id": newFile.ID,
	})
}
