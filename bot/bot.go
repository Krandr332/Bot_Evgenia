package bot

import (
	"log"
	"sync"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *tgbotapi.BotAPI
var wg sync.WaitGroup

func StartBot() {
	// Инициализация Telegram бота
	botToken := "YOUR_BOT_TOKEN_HERE"
	var err error
	bot, err = tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Не удалось инициализировать бота:", err)
	}

	// Другой код для обработки команд и действий бота
	// ...

	// Ожидание завершения всех горутин перед выходом
	wg.Wait()
}
