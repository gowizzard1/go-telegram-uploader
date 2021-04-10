package main

import (
	"fmt"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("haha")
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
