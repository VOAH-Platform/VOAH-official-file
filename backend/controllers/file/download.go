package file

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/models"
	"implude.kr/VOAH-Official-File/utils/fileutil"
)

func DownloadFileCtrl(c *fiber.Ctx) error {
	fileID, err := strconv.Atoi(c.Query("file-id", ""))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   "File ID is Integer",
		})
	}
	db := database.DB
	var fileMetadata *models.File
	err = db.Where(&models.File{ID: uint(fileID)}).First(&fileMetadata).Error
	if err == gorm.ErrRecordNotFound {
		return c.Status(404).JSON(fiber.Map{
			"message": "File Not Exists or Deleted",
		})
	} else if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	if fileMetadata.Processed == 2 {
		c.Set("Content-Type", fileMetadata.FileType)
		if configs.Env.File.USE_S3 {
			fileByte, err := fileutil.ReadFileFromS3(fileID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
			return c.Send(fileByte)

		} else {
			// new io.Reader
			fileReader, err := fileutil.ReadFileFromLocal(fileID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
			return c.SendStream(fileReader)
		}
	} else {
		return c.Status(400).JSON(fiber.Map{
			"message": "File is not processed or failed to process",
		})
	}
}
