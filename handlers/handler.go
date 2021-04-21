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
	Name string
	Path string
	Dir  bool
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

	path := "files/"

	if len(c.Params("*")) > 0 {
		path += c.Params("*") + "/"
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
			Path: path + file.Name(),
		})
	}

	return c.Render("files", fiber.Map{
		"Files": vd.Files,
	})
}

func SendFile(c *fiber.Ctx) error {
	return fmt.Errorf("s")
}

func Ex(c *fiber.Ctx) error {
	return c.Render("ex", fiber.Map{
		"Result": tools.GetRandomString(12),
	})
}
