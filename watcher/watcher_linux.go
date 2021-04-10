package watcher

import (
	"powergen/go-telegram-uploader/config"

	"github.com/eloylp/go-telegram-uploader/fails"
	"github.com/eloylp/go-telegram-uploader/handler"
	"github.com/rjeczalik/notify"
)

func Watcher(conf config.Configs) {
	events := make(chan notify.EventInfo, 1)
	err := notify.Watch(conf.FolderToScan, events, notify.InCloseWrite)
	fails.FailIfError(err)
	startWatcher(events, handler.ProcessFile)
}

func startWatcher(events chan notify.EventInfo, handler func(string)) {

	for event := range events {
		if event.Event() == notify.InCloseWrite {
			go handler(event.Path())
		}
	}
}
