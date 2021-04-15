package main

import (
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("index", fiber.Map{
			"Title": getRandomString(10),
			"Some":  getRandomString(15),
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func getRandomString(strLen int) string {

	alph := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, strLen)
	for i := range s {
		s[i] = alph[rand.Intn(len(alph))]
	}

	return string(s)
}
