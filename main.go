package main

import (
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
)

func main() {

	color.Info.Tips("Telegram video uploader started -- drop or copy to /home directory")
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
