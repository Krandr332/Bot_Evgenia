package bot

import (
	"log"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
)

// Создаем структуру для бота
type Bot struct {
	api    *tgbotapi.BotAPI
	update tgbotapi.Update
}

// Инициализируем бота
func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{api: api}, nil
}

// Запускаем бота
func (b *Bot) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			if checkForUserInSystem(fmt.Sprintf("%d", update.Message.From.ID)) > 0 {
				fmt.Println("ты в системе уже есть")
			}else{
				b.handleStart(update.Message.Chat.ID)
			}
			fmt.Println( update.Message.From.ID)
			
		}
		
	}
}

// Обработка команды /start
func (b *Bot) handleStart(chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "Привет! Нажмите кнопку 'Начать' на клавиатуре ниже:")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("/start"),
		),
	)

	_, err := b.api.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
