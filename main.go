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

	api.Post("/random_string", handlers.GetRandomResult)
	api.Get("/ex", handlers.Ex)
	api.Post("/files/*", handlers.GetFiles)

	app.Get("/", handlers.GetMain)
	app.Get("/home/:user?", handlers.GetUserFromURL)
	app.Get("/random.txt", handlers.GetRandomTXT)
	app.Get("/ex", handlers.Ex)

	app.Post("/send", handlers.Send)

	log.Fatal(app.Listen(":3000"))
}
