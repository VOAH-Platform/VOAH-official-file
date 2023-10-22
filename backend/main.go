package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"implude.kr/VOAH-Official-File/configs"
	"implude.kr/VOAH-Official-File/database"
	"implude.kr/VOAH-Official-File/routers"
	"implude.kr/VOAH-Official-File/utils/directory"
	"implude.kr/VOAH-Official-File/utils/logger"
)

func main() {
	configs.LoadEnv()   // Load envs
	logger.InitLogger() // Intitialize logger
	directory.IniteDirectory()
	database.ConnectDB()

	serverConf := configs.Env.Server
	log := logger.Logger

	app := fiber.New(fiber.Config{
		BodyLimit: configs.Env.File.FileSizeLimitMB * 1024 * 1024, // this is the default limit of 1GB
	})

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: serverConf.CSRFOrigin,
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	routers.Initialize(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConf.Port)))
}
