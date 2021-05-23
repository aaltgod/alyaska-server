package main

import (
	"log"

	handler "github.com/alyaskastorm/fiber_example/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{
		BodyLimit: 250 * 1024 * 1024,
	})

	app.Use(logger.New())
	app.Use(cors.New())

	api := app.Group("/api")

	api.Post("/files", handler.GetFiles)
	api.Post("/get-file", handler.SendFile)
	api.Post("/upload-file", handler.UploadFile)
	api.Get("/folder/:folderName", handler.GetFolder)

	log.Fatal(app.Listen(":3000"))
}
