package handler

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"powergen/go-telegram-uploader/config"
	"regexp"

	"powergen/go-telegram-uploader/bot"

	"powergen/go-telegram-uploader/fails"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ProcessFile(path string) {

	videoConfig, err := checkForVideo(path)
	if err == nil {
		bot.SendToBot(videoConfig)
		return
	}
}

func checkForVideo(path string) (tgbotapi.VideoConfig, error) {
	fileName := filepath.Base(path)
	matched, err := regexp.MatchString(`(?i).*\.(avi|mp4|flv|mov|mkv)`, fileName)
	fails.FailIfError(err)
	conf := config.BuildConfigs()
	if matched {
		data, err := ioutil.ReadFile(path)
		fails.FailIfError(err)
		photoData := tgbotapi.FileBytes{Name: fileName, Bytes: data}
		dat := tgbotapi.NewVideoUpload(conf.ChatId, photoData)
		dat.Duration = 2
		return dat, nil
	}
	return tgbotapi.VideoConfig{}, errors.New("cannot map for video")
}
