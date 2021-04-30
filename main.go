package main

import (
	"log"

	"github.com/alyaskastorm/fiber_example/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	app.Use(logger.New())

	api := app.Group("/api")

	api.Post("/random-string", handlers.GetRandomResult)
	api.Post("/files", handlers.GetFiles)
	api.Post("/get-file", handlers.SendFile)

	log.Fatal(app.Listen(":3000"))
}
