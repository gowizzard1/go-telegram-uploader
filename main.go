package main

import (
	"fmt"
	"os"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
)

func main() {
	os.Exit(500)
	fmt.Println()
	fmt.Println()
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
