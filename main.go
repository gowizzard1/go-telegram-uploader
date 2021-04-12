package main

import (
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
)

func main() {

	color.Info.Tips("starting the go uploader app")
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
