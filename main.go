package main

import (
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
