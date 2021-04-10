package move

import (
	"log"
	"os"
	"path/filepath"
	"powergen/go-telegram-uploader/config"
)

func MoveFile(conf config.Configs, pathName string) {
	file := filepath.Base(pathName)
	oldLocation := conf.FolderToScan + "/" + file
	newLocation := conf.FolderToMove + "/" + file
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
