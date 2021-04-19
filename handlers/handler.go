package handlers

import (
	"crypto/sha1"
	"fmt"

	"github.com/alyaskastorm/fiber_example/tools"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	Str1, Str2 string
}

type File struct {
	Name string
}

type ViewData struct {
	Items []Item
	Files []File
}

func Send(c *fiber.Ctx) error {

	msg := c.FormValue("msg")
	hash := sha1.New()

	hash.Write([]byte(msg))

	result := fmt.Sprintf("%x", hash.Sum(nil))

	return c.Render("result", fiber.Map{
		"Result": result,
	})
}

func GetUserFromURL(c *fiber.Ctx) error {
	return c.Render("user", fiber.Map{
		"Name": c.Params("user"),
	})
}

func GetRandomTXT(c *fiber.Ctx) error {
	return c.Download("./files/random.txt")
}

func GetMain(c *fiber.Ctx) error {

	d := new(ViewData)
	d.Items = []Item{
		{
			Str1: tools.GetRandomString(2),
			Str2: tools.GetRandomString(4)},
		{
			Str1: tools.GetRandomString(2),
			Str2: tools.GetRandomString(4),
		}}

	return c.Render("index", fiber.Map{
		"Title": tools.GetRandomString(10),
		"Some":  tools.GetRandomString(15),
		"Items": d.Items,
	})
}
