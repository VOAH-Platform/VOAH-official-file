package routers

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Official-File/controllers/file"
)

func addFile(router *fiber.App) {
	fileGroup := router.Group("/api/file") // auth router

	fileGroup.Get("/", func(c *fiber.Ctx) error {
		return file.CheckFileCtrl(c)
	})
	fileGroup.Post("/", func(c *fiber.Ctx) error {
		return file.UploadFileCtrl(c)
	})
	fileGroup.Delete("/", func(c *fiber.Ctx) error {
		return file.DeleteFileCtrl(c)
	})

	fileGroup.Get("/download", func(c *fiber.Ctx) error {
		return file.DownloadFileCtrl(c)
	})
}
