package main

import (
	// "fmt"
	"log"
	// "os"
	// "time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Инициализация бота с токеном
	bot, err := tgbotapi.NewBotAPI("5435221086:AAFtVvXSL4ZJ427yHpCQXkqVnjUz0eHI3C8")
	if err != nil {
		log.Panic(err)
	}

	// Установка опций для работы в асинхронном режиме
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// Если пришло новое сообщение
		if update.Message == nil {
			continue
		}

		// Отправляем "Привет, мир!" при получении команды /start
		if update.Message.IsCommand() {
			command := update.Message.Command()
			switch command {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, мир!")
				bot.Send(msg)
			}
		} else {
			// Отвечаем на обычные текстовые сообщения
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Получено сообщение: "+update.Message.Text)
			bot.Send(msg)
		}
	}
}
