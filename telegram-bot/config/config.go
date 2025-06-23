package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	TelegramToken string
	BackendURL    = "http://localhost:8081"
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env файл не найден, читаю переменные окружения напрямую")
	}

	TelegramToken = os.Getenv("TELEGRAM_TOKEN")
}
