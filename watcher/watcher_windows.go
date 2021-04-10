package watcher

import (
	"log"

	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/fails"

	"powergen/go-telegram-uploader/handler"

	"github.com/fsnotify/fsnotify"
)

func Watcher(conf config.Configs) {
	createdWatcher, err := fsnotify.NewWatcher()
	fails.FailIfError(err)
	err = createdWatcher.Add(conf.FolderToScan)
	fails.FailIfError(err)
	startWatcher(createdWatcher, handler.ProcessFile)
}

func startWatcher(watcher *fsnotify.Watcher, handler func(string)) {

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				filePath := event.Name
				go handler(filePath)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			if err != nil {
				log.Println("error:", err)
				return
			}

		}
	}
}
