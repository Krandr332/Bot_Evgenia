package bot

import (
	"fmt"
	"log"

	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	api        *tgbotapi.BotAPI
	userStates map[int64]*UserState
}

type UserState struct {
	State       RegistrationState
	FullName    string
	PhoneNumber string
	Email       string
}

type RegistrationState int

const (
	StateStart RegistrationState = iota
	StateFullName
	StatePhoneNumber
	StateEmail
	StateComplete
)

func NewBot(token string) (*Bot, error) {
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &Bot{
		api:        api,
		userStates: make(map[int64]*UserState),
	}, nil
}

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

		chatID := update.Message.Chat.ID
		userState, ok := b.userStates[chatID]

		if !ok {
			userState = &UserState{State: StateStart}
			b.userStates[chatID] = userState
		}

		if userState.State == StateStart {
			b.handleStart(chatID, userState)
		} else {
			b.handleRegistration(update.Message, userState)
		}
	}
}

func (b *Bot) handleStart(chatID int64, userState *UserState) {
	if userState.State == StateStart {
		msg := tgbotapi.NewMessage(chatID, "Доброго времени суток! Вы не зарегистрированы, давайте знакомиться. Пожалуйста, отправьте свое ФИО.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
		userState.State = StateFullName
	}
}

func (b *Bot) handleRegistration(message *tgbotapi.Message, userState *UserState) {
	chatID := message.Chat.ID

	switch userState.State {
	case StateFullName:
		userState.FullName = message.Text
		userState.State = StatePhoneNumber

		msg := tgbotapi.NewMessage(chatID, "Отправьте ваш номер телефона.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	case StatePhoneNumber:
		userState.PhoneNumber = message.Text
		userState.State = StateEmail

		msg := tgbotapi.NewMessage(chatID, "Отправьте вашу почту.")
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	case StateEmail:
		userState.Email = message.Text
		userState.State = StateComplete

		msg := tgbotapi.NewMessage(chatID, "Регистрация завершена. Спасибо!")
		fmt.Println(userState)
		_, err := b.api.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}
}
