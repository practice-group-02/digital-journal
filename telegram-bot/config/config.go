package config

import "os"

func GetTelegramToken() string {
	return os.Getenv("TELEGRAM_TOKEN")
}
