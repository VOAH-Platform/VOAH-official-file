package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"implude.kr/VOAH-Template-Project/configs"
	"implude.kr/VOAH-Template-Project/routers"
	"implude.kr/VOAH-Template-Project/utils/logger"
)

func main() {
	configs.LoadEnv()   // Load envs
	logger.InitLogger() // Intitialize logger

	serverConf := configs.Env.Server
	log := logger.Logger

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: serverConf.CSRFOrigin,
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	routers.Initialize(app)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", serverConf.Port)))
}
