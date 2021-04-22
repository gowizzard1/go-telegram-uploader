package watcher

import (
	"fmt"
	"powergen/go-telegram-uploader/config"

	"powergen/go-telegram-uploader/handler"

	"github.com/gookit/color"
	"github.com/rjeczalik/notify"
)

func Watcher(conf config.Configs) {
	events := make(chan notify.EventInfo, 1)
	err := notify.Watch(conf.FolderToScan, events, notify.InCloseWrite)
	if err != nil {
		color.Error.Println(err.Error())
	}
	startWatcher(events, handler.ProcessFile)
}

func startWatcher(events chan notify.EventInfo, handler func(string)) {

	for event := range events {
		if event.Event() == notify.InCloseWrite {
			fmt.Println("handling the video")
			go handler(event.Path())
		}
	}
}
