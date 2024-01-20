package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v3"
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

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/hello", func(c tele.Context) error {
		return c.Send("Hello!")
	})

	b.Handle(tele.OnText, func(c tele.Context) error {
		// All the text messages that weren't
		// captured by existing handlers.

		var (
			user = c.Sender()
			text = c.Text()
		)

		// Use full-fledged bot's functions
		// only if you need a result:
		msg, err := b.Send(user, text)
		if err != nil {
			return err
		}

		// Instead, prefer a context short-hand:
		c.Send("OnText / msg.Payload: " + msg.Payload)
		return c.Send(text)
	})

	b.Handle(tele.OnChannelPost, func(c tele.Context) error {
		// Channel posts only.
		msg := c.Message()
		return c.Send("OnChannelPost: " + msg.Text)
	})

	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		// Photos only.
		photo := c.Message().Photo
		return c.Send(photo)
	})

	// b.Handle(tele.OnQuery, func(c tele.Context) error {
	// 	// Incoming inline queries.
	// 	return c.Answer("OnQuery")
	// })

	b.Start()
}
