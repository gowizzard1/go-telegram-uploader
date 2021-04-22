package move

import (
	"log"
	"os"
	"path/filepath"
	"powergen/go-telegram-uploader/config"
	"strings"
)

func MoveFile(conf config.Configs, pathName string) {
	file := filepath.Base(pathName)
	oldLocation := RemoveLastPart(conf.FolderToScan + "/" + file)
	newLocation := RemoveLastPart(conf.FolderToMove + file) //file has .filepart before finish moving
	if _, err := os.Stat(newLocation); os.IsNotExist(err) {
		os.Mkdir(newLocation, os.ModeDir)
	}
	er := os.Rename(oldLocation, newLocation)
	if er != nil {
		log.Fatal(er)
	}
}

func RemoveLastPart(path string) string {
	ss := strings.Split(path, ".")
	st := ss[len(ss)-3]
	su := ss[len(ss)-2]
	return st + "." + su
}
