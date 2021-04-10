package main

import (
	"fmt"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
	log "github.com/visionmedia/go-cli-log"
)

func main() {
	fmt.Println()
	log.Info("unpack", "tarball to node_modules/express")
	fmt.Println()
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
