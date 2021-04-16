package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

type Item struct {
	Str1, Str2 string
}

type ViewData struct {
	Items []Item
}

func main() {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Post("/send", func(c *fiber.Ctx) error {
		msg := c.FormValue("msg")
		h := sha1.New()

		h.Write([]byte(msg))

		result := fmt.Sprintf("%x", h.Sum(nil))

		return c.Render("result", fiber.Map{
			"Result": result,
		})
	})

	app.Get("/home/:user?", func(c *fiber.Ctx) error {
		return c.Render("user", fiber.Map{
			"Name": c.Params("user"),
		})
	})

	app.Get("/random.txt", func(c *fiber.Ctx) error {
		return c.Download("./files/random.txt")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		d := new(ViewData)
		d.Items = []Item{
			{
				Str1: getRandomString(2),
				Str2: getRandomString(4)},
			{
				Str1: getRandomString(2),
				Str2: getRandomString(4),
			}}
		return c.Render("index", fiber.Map{
			"Title": getRandomString(10),
			"Some":  getRandomString(15),
			"Items": d.Items,
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
