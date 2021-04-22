package main

import (
	"fmt"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/gookit/color"
	"github.com/joho/godotenv"
)

func main() {
	color.Info.Tips("Telegram video uploader starting")
	godotenv.Load()
	conf := config.BuildConfigs()
	fmt.Println("Telegram video uploader started >> drop or copy to ", conf.FolderToScan)
	watcher.Watcher(conf)

}
