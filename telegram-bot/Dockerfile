# Используем официальный образ Go
FROM golang:1.24-alpine

# Рабочая директория
WORKDIR /app

# Копируем модули Go
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN go build -o bot .

ENV TELEGRAM_TOKEN=8168169856:AAEh6I7uQyl9dKale7E21uoEiVWXzsQAsFE

# Запускаем бота
CMD ["./bot"]
