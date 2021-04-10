package main

import (
	"log"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Starting app. This will be terminated by CR+LF\r")
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
