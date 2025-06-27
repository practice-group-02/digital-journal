package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const apiUrl = "http://backend:8080"

// Переменная для хранения ID чата
var chatID int64

var bot *tgbotapi.BotAPI

// Инициализация бота
func init() {
	var err error
	bot, err = tgbotapi.NewBotAPI("8168169856:AAEh6I7uQyl9dKale7E21uoEiVWXzsQAsFE")
	if err != nil {
		log.Fatal(err)
	}
}

// Функция для отправки сообщений пользователю
func sendMessage(msg string) {
	if len(msg) > 4096 {
		// Разбиваем сообщение на несколько частей, если оно слишком длинное
		for i := 0; i < len(msg); i += 4096 {
			part := msg[i:min(i+4096, len(msg))]
			message := tgbotapi.NewMessage(chatID, part)
			bot.Send(message)
		}
	} else {
		// Отправляем сообщение в одном запросе
		message := tgbotapi.NewMessage(chatID, msg)
		bot.Send(message)
	}
}

// Функция для получения программ с сервера
func getPrograms() string {
	resp, err := http.Get(apiUrl + "/programs")
	if err != nil {
		log.Printf("Error fetching programs: %s", err)
		return "Error fetching programs"
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: Received non-OK status: %d", resp.StatusCode)
		return fmt.Sprintf("Error fetching programs, status code: %d", resp.StatusCode)
	}

	var programs []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&programs); err != nil {
		log.Printf("Error decoding response: %s", err)
		return "Error decoding response"
	}

	// Если нет программ, возвращаем сообщение
	if len(programs) == 0 {
		return "No programs found."
	}

	// Форматируем ответ
	var result strings.Builder
	for _, program := range programs {
		result.WriteString(fmt.Sprintf("Title: %s\nType: %s\nDescription: %s\n\n", program["title"], program["type"], program["description"]))
	}
	return result.String()
}

func getProgramsByType(programType string) string {
	resp, err := http.Get(apiUrl + "/programs/" + programType)
	if err != nil {
		log.Printf("Error fetching programs by type: %s", err)
		return "Error fetching programs by type"
	}
	defer resp.Body.Close()

	var programs []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&programs); err != nil {
		return "Error decoding response"
	}

	var result strings.Builder
	for _, program := range programs {
		result.WriteString(fmt.Sprintf("Title: %s\nType: %s\nDescription: %s\n\n", program["title"], program["type"], program["description"]))
	}
	return result.String()
}

// // Функция для получения программ по тегам
// func getProgramsByTags(tags string) string {
// 	// Преобразуем теги в нижний регистр, если это необходимо
// 	tags = strings.ToLower(tags)

// 	// Формируем запрос к серверу
// 	resp, err := http.Get(apiUrl + "/programs/tags?tags=" + tags)
// 	if err != nil {
// 		log.Printf("Error fetching programs by tags: %s", err)
// 		return "Error fetching programs by tags"
// 	}
// 	defer resp.Body.Close()

// 	// Чтение ответа
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Printf("Error reading response body: %s", err)
// 		return "Error reading response body"
// 	}

// 	// Декодируем JSON в структуру Program
// 	var programs []models.Program // Используем структуру, которая уже есть на бэкенде
// 	if err := json.Unmarshal(body, &programs); err != nil {
// 		log.Printf("Error decoding response: %s", err)
// 		return "Error decoding response"
// 	}

// 	// Если программ нет, возвращаем сообщение
// 	if len(programs) == 0 {
// 		return "No programs found for the specified tags."
// 	}

// 	// Форматируем ответ для Telegram
// 	var result strings.Builder
// 	for _, program := range programs {
// 		result.WriteString(fmt.Sprintf("Title: %s\nType: %s\nDescription: %s\n", program.Title, program.Type, program.Description))
// 		result.WriteString("Tags: ")
// 		for _, tag := range program.Tags {
// 			result.WriteString(fmt.Sprintf("%s ", tag.Name))
// 		}
// 		result.WriteString("\n\n")
// 	}

// 	return result.String() // Возвращаем отформатированное сообщение
// }

// Функция для обработки входящих сообщений
func handleMessages(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	chatID = update.Message.Chat.ID
	log.Printf("Received message: %s", update.Message.Text) // Логирование входящего сообщения

	// Обработка команд
	switch {
	// Обрабатываем команду "/programs/<type>"
	case strings.HasPrefix(update.Message.Text, "/programs/"):
		// Извлекаем тип программы из команды
		programType := strings.TrimPrefix(update.Message.Text, "/programs/")
		programType = strings.TrimSpace(programType)

		// Проверяем, если тип программы не указан
		if programType == "" {
			sendMessage("Please specify the type of program, e.g., /programs/стипендия")
			return
		}

		// Отправляем запрос на сервер для получения программ по типу
		sendMessage("Fetching programs of type: " + programType)
		sendMessage(getProgramsByType(programType)) // Отправляем запрос по типу программы

	// // Обрабатываем команду "/programs/tags?tags=<tags>"
	// case strings.HasPrefix(update.Message.Text, "/programs/tags"):
	// 	// Извлекаем теги из команды
	// 	tags := strings.TrimPrefix(update.Message.Text, "/programs/tags")
	// 	tags = strings.TrimSpace(tags)

	// 	// Проверяем, если теги не указаны
	// 	if tags == "" {
	// 		sendMessage("Please provide tags, e.g., /programs/tags магистратура,научные исследования")
	// 		return
	// 	}

	// 	// Отправляем запрос на сервер для получения программ по тегам
	// 	sendMessage("Fetching programs with tags: " + tags)
	// 	programsMessage := getProgramsByTags(tags) // Отправляем запрос по тегам
	// 	sendMessage(programsMessage)               // Отправляем найденные программы

	// Обработка команды по умолчанию для команды "/programs"
	case update.Message.Text == "/programs":
		sendMessage("Fetching all programs...")
		time.Sleep(time.Duration(5) * time.Second) // Добавлен таймаут перед отправкой
		programs := getPrograms()                  // Получаем все программы
		log.Printf("Programs: %s", programs)       // Логирование полученных программ
		sendMessage(programs)                      // Отправка сообщения в Telegram

	default:
		sendMessage("Invalid command. Please use /programs/<type> (e.g., /programs/стипендия) or /programs/tags?tags=<tag1>,<tag2>.")
	}
}

// Функция для нахождения минимального значения
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	log.Printf("TG WORKS")

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		handleMessages(update)
	}
}
