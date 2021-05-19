package handlers

import (
	"log"
	"os"
	"strings"

	"github.com/alyaskastorm/fiber_example/tools"
	"github.com/gofiber/fiber/v2"
)

type File struct {
	Name         string `json:"name"`
	Path         string `json:"path,omitempty"`
	PreviousPath string `json:"previous_path,omitempty"`
	Dir          bool   `json:"dir,omitempty"`
}

func GetFiles(c *fiber.Ctx) error {

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
	defer file.Close()

	dirs, err := file.Readdir(-1)
	if err != nil {
		log.Println(err)
	}

	dirs = tools.SortDir(dirs)

	var prevPath string

	p := strings.Split(path, "/")
	if len(p) > 1 {
		prevPath = strings.Join(p[:len(p)-1], "/")
	}

	files := []File{}

	for _, file := range dirs {
		files = append(files, File{
			Name:         file.Name(),
			Dir:          file.IsDir(),
			Path:         path + "/" + file.Name(),
			PreviousPath: prevPath,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"files": files,
	})
}

func SendFile(c *fiber.Ctx) error {

	f := new(File)

	if err := c.BodyParser(f); err != nil {
		log.Println(err)

		return err
	}

	return c.Download(f.Path)
}

func UploadFile(c *fiber.Ctx) error {

	multiForm, err := c.MultipartForm()
	if err != nil {
		log.Println(err)

		return err
	}

	folderName := tools.GetRandomString(6)
	if err := os.Mkdir("uploads/"+folderName, 0755); err != nil {
		log.Println(err)

		return err
	}

	for _, file := range multiForm.File {
		if err := c.SaveFile(file[0], "uploads/"+folderName+"/"+file[0].Filename); err != nil {
			log.Println(err)

			return err
		}

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"folderName": folderName,
		"folderPath": "uploads/" + folderName,
	})
}

func GetFolder(c *fiber.Ctx) error {

	folderName := c.Params("folderName")

	file, err := os.Open("uploads/" + folderName)
	if err != nil {
		log.Println(err)

		return err
	}
	defer file.Close()

	dirs, err := file.Readdir(-1)
	if err != nil {
		log.Println(err)

		return err
	}

	files := []File{}

	for _, file := range dirs {
		files = append(files, File{
			Name: file.Name(),
			Path: "uploads/" + folderName + "/" + file.Name(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"files": files,
	})
}
