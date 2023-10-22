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

func DeleteFileCtrl(c *fiber.Ctx) error {
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
		if configs.Env.File.USE_S3 {
			err = fileutil.DeleteFileFromS3(fileID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
		} else {
			err = fileutil.DeleteFileFromLocal(fileID)
			if err != nil {
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal Server Error",
				})
			}
		}
	} else if fileMetadata.Processed == 1 {
		return c.Status(400).JSON(fiber.Map{
			"message": "File is not processed",
		})
	}

	err = db.Delete(&fileMetadata).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
