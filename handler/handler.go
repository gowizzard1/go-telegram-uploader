package handler

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"powergen/go-telegram-uploader/bot"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/move"
	"powergen/go-telegram-uploader/validation"
	"regexp"

	"powergen/go-telegram-uploader/fails"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func ProcessFile(path string) {
	conf := config.BuildConfigs()
	videoConfig, err := checkForVideo(path)
	if err == nil {
		fmt.Println("sending video...")
		bot.SendToBot(videoConfig)
		fmt.Println("moving sent video...")
		move.MoveFile(conf, path)
		return
	}
	fmt.Println(err)
}

func checkForVideo(path string) (vc tgbotapi.VideoConfig, err error) {
	fileName := filepath.Base(path)
	matched, err := regexp.MatchString(`(?i).*\.(avi|mp4|flv|mov|mkv)`, fileName)
	fails.FailIfError(err)
	conf := config.BuildConfigs()
	if matched {
		//check size
		info := validation.GetVideoInfo(path)
		fmt.Println("video size is ", validation.LenReadable(int(info.Size()), 2))
		fileSize := info.Size() / 250000000
		if fileSize > 2 {
			return vc, errors.New("the video size is greater than 2GB")
		}
		//read the file
		data, err := ioutil.ReadFile(path)
		fails.FailIfError(err)
		photoData := tgbotapi.FileBytes{Name: fileName, Bytes: data}
		dat := tgbotapi.NewVideoUpload(conf.ChatId, photoData)
		dat.Duration = 2
		return dat, nil
	}
	return tgbotapi.VideoConfig{}, errors.New("cannot map for video")
}
