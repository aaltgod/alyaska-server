package tools

import (
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func QRcode(folderName string) error {

	err := qrcode.WriteFile(
		"http://185.252.146.12:8080/folder/"+folderName,
		qrcode.Medium,
		256,
		"frontend/src/assets/qrcodes/"+folderName+".png",
	)
	if err != nil {
		log.Println(err)

		return err
	}

	return nil
}
