package main

import (
	"log"

	handler "github.com/alyaskastorm/fiber_example/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New(fiber.Config{})

	app.Use(cors.New())

	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/random-string", handler.GetRandomResult)
	api.Post("/files", handler.GetFiles)
	api.Post("/get-file", handler.SendFile)
	api.Post("/upload-file", handler.UploadFile)

	log.Fatal(app.Listen(":3000"))
}
