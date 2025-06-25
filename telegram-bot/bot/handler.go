package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	switch update.Message.Text {
	case "/start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, `🎓 Добро пожаловать в Scholar.KZ!

Здесь ты найдёшь актуальные 🇰🇿 гранты, программы обучения и дедлайны.
Вот что я умею:

📌 /latest — последние гранты  
🔍 /search — фильтр по стране  
📅 /deadline — ближайшие дедлайны

Начни с команды /latest, чтобы увидеть примеры.
`)
		bot.Send(msg)

	case "/latest":
		programs, err := GetLatestPrograms()
		if err != nil {
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при получении данных."))
			return
		}

		for _, p := range programs {
			text := fmt.Sprintf("🎓 %s (%s)\n📅 Дедлайн: %s\n%s", p.Title, p.Country, p.Deadline, p.Description)
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, text))
		}

	default:
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда 😕"))
	}
}
