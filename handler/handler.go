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
	Path         string `json:"path"`
	PreviousPath string `json:"previous_path"`
	Dir          bool   `json:"dir"`
}

type ViewData struct {
	Files []File
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

	vd := new(ViewData)

	for _, file := range dirs {
		vd.Files = append(vd.Files, File{
			Name:         file.Name(),
			Dir:          file.IsDir(),
			Path:         path + "/" + file.Name(),
			PreviousPath: prevPath,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"files": vd.Files,
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

func GetRandomResult(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"string":  tools.GetRandomString(10),
	})
}

func UploadFile(c *fiber.Ctx) error {

	multiForm, err := c.MultipartForm()
	if err != nil {
		log.Println(err)

		return err
	}

	folderName := tools.GetRandomString(10)
	if err := os.Mkdir("uploads/"+folderName, 0755); err != nil {
		log.Println(err)

		return err
	}

	for key, file := range multiForm.File {

		if err := c.SaveFile(file[0], "uploads/"+folderName+"/"+"file"+key); err != nil {
			log.Println(err)

			return err
		}
	}

	return nil
}
