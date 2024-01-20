package main

import (
	"log"
	"fmt"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке файла .env")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN не найден в файле .env")
	}

	chatIDstring := os.Getenv("TELEGRAM_CHAT_ID")
	if chatIDstring == "" {
		log.Fatal("TELEGRAM_CHAT_ID не найден в файле .env")
	}

	var chatID int64
	_, err := fmt.Sscanf(chatIDstring, "%d", &chatID)
	if err != nil {
		log.Fatal("Ошибка при преобразовании TELEGRAM_CHAT_ID в int64")
	}

	// creates a new BotAPI instance
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	message := tgbotapi.NewMessage(chatID, "Привет, мир! Это ваш бот на Golang.")
	_, err = bot.Send(message)
	if err != nil {
		log.Panic(err)
	}
}
