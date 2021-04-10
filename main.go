package main

import (
	"fmt"
	"log"
	"os/exec"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"

	"github.com/joho/godotenv"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
