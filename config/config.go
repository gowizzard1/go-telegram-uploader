package config

import (
	"os"
)

type Configs struct {
	BotToken     string `env:"BOT_TOKEN,required"`
	ChatId       int64  `env:"CHAT_ID,required"`
	FolderToScan string `env:"FOLDER_TO_SCAN,required"`
	FolderToMove string `env:"FOLDER_TO_MOVE"`
}

var configs Configs

func SetUp(conf Configs) {
	configs = conf
}

func GetConfigs() Configs {
	return configs
}

// buildConfigs reads the environment variables and builds the configurations
func BuildConfigs() Configs {
	var conf Configs

	conf = Configs{
		ChatId:       -525034720,
		FolderToScan: getEnv("FOLDER_TO_SCAN", ""),
		BotToken:     getEnv("BOT_TOKEN", "key"),
		FolderToMove: getEnv("FOLDER_TO_MOVE", ""),
	}
	SetUp(conf)
	return conf
}

// getEnv is a helper function to read an environment variable or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
