package main

import (
	"fmt"
	"os"
	"os/exec"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/watcher"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	cmd := exec.Command("pass")
	cmd.Stdin = strings.NewReader("Secret pass!\n")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
	fmt.Println()
	fmt.Println()
	godotenv.Load()
	conf := config.BuildConfigs()
	watcher.Watcher(conf)
}
