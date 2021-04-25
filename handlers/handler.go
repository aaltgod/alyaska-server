package handlers

import (
	"crypto/sha1"
	"fmt"
	"log"
	"os"

	"github.com/alyaskastorm/fiber_example/tools"
	"github.com/gofiber/fiber/v2"
)

type Item struct {
	Str1, Str2 string
}

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Dir  bool   `json:"dir"`
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

func GetFiles(c *fiber.Ctx) error {
	log.Println(c.Request().String())

	f := File{}

	if err := c.BodyParser(&f); err != nil {
		log.Println(err)

		return err
	}

	path := f.Path
	if len(path) == 0 {
		path = "files"
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println(err)

		return err
	}

	dirs, err := file.Readdir(-1)
	if err != nil {
		log.Println(err)
	}

	vd := new(ViewData)

	for _, file := range dirs {
		vd.Files = append(vd.Files, File{
			Name: file.Name(),
			Dir:  file.IsDir(),
			Path: path + "/" + file.Name(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"files": vd.Files,
	})
}

func SendFile(c *fiber.Ctx) error {
	return fmt.Errorf("s")
}

func Ex(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "application/json; charset=UTF-8")
	c.Response().Header.Add("Access-Control-Allow-Origin", "*")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ex": tools.GetRandomString(12),
	})
}

func GetRandomResult(c *fiber.Ctx) error {
	c.Response().Header.Add("Content-Type", "application/json; charset=UTF-8")
	c.Response().Header.Add("Access-Control-Allow-Origin", "*")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"string":  tools.GetRandomString(10),
	})
}
