package file

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/models"
)

func CheckFileCtrl(c *fiber.Ctx) error {
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

	return c.JSON(fiber.Map{
		"message": "success",
		"file":    fileMetadata,
	})

}
