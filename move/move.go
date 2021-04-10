package move

import (
	"log"
	"os"
	"powergen/go-telegram-uploader/config"
)

func MoveFile(ext string) {
	conf := config.BuildConfigs()
	oldLocation := conf.FolderToScan + ext
	newLocation := conf.FolderToMove + ext
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
