package routers

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Template-Project/controllers/info"
)

func addInfo(router *fiber.App) {
	infoGroup := router.Group("/api/info") // auth router

	infoGroup.Get("", func(c *fiber.Ctx) error {
		return info.GetInfoCtrl(c)
	})
}
