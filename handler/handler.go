package handler

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"powergen/go-telegram-uploader/bot"
	"powergen/go-telegram-uploader/config"
	"powergen/go-telegram-uploader/move"
	"powergen/go-telegram-uploader/validation"
	"regexp"

	"powergen/go-telegram-uploader/fails"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gookit/color"
)

func ProcessFile(path string) {
	conf := config.BuildConfigs()
	videoConfig, err := checkForVideo(path)
	if err == nil {
		color.Info.Tips("sending video" + filepath.Base(path))
		bot.SendToBot(videoConfig)
		color.Green.Println("video " + filepath.Base(path) + " sent")
		color.Info.Tips("preparing to move sent " + filepath.Base(path) + " to " + conf.FolderToMove)
		move.MoveFile(conf, path)
		color.Green.Println("video " + filepath.Base(path) + " moved")
		return
	}
	if err.Error() == "the video size is greater than 2GB" {
		color.Warn.Println("the video size is greater than 2GB")
	}
}

func checkForVideo(path string) (vc tgbotapi.VideoConfig, err error) {
	fileName := filepath.Base(path)
	matched, err := regexp.MatchString(`(?i).*\.(avi|mp4|flv|mov|mkv)`, fileName)
	fails.FailIfError(err)
	conf := config.BuildConfigs()
	if matched {
		//check size
		info := validation.GetVideoInfo(path)
		color.Yellow.Println("video size is ", validation.LenReadable(int(info.Size()), 2))
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
