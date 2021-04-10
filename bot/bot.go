package bot

import (
	"log"
	"sync"

	"powergen/go-telegram-uploader/config"

	"powergen/go-telegram-uploader/fails"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI
var once sync.Once

func GetBot() *tgbotapi.BotAPI {
	once.Do(func() {
		conf := config.BuildConfigs()
		telegramToken := conf.BotToken
		createdBot, err := tgbotapi.NewBotAPI(telegramToken)
		fails.FailIfError(err)
		bot = createdBot
	})
	return bot
}

func SendToBot(chattable tgbotapi.Chattable) {
	_, err := GetBot().Send(chattable)
	log.Println(err)
}
