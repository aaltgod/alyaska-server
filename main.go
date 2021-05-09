package main

import (
	"log"

	h "github.com/alyaskastorm/fiber_example/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(cors.New())

	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/random-string", h.GetRandomResult)
	api.Post("/files", h.GetFiles)
	api.Post("/get-file", h.SendFile)
	api.Post("/upload-file", h.UploadFile)
	api.Get("/folder/:folderName", h.GetFolder)

	log.Fatal(app.Listen(":3000"))
}
