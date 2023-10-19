package bot

import (
	"fmt"
	"log"
	"strings"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

type Bot struct {
	api        *tgbotapi.BotAPI
	userStates map[int64]*UserState
}

type UserState struct {
	State       RegistrationState
	
	FirstName   string
	LastName    string
	MiddleName  string
	PhoneNumber string
	Email       string
	Region      string
}

type RegistrationState int

const (
	StateStart RegistrationState = iota
	StateFullName
	StateRegion
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
        tgID := update.Message.From.ID

        if checkForUserInSystem(fmt.Sprintf("%d", update.Message.From.ID)) == 0 {
			// Проверка на зарегес трированность 
			if userState.State == StateStart {
				b.handleStart(chatID, userState)
			} else {
				b.handleRegistration(update.Message, userState)
			}
        }else{
			isAdmin, err := checkAdminStatus(tgID)
        if err != nil {
            log.Println("Ошибка при проверке admin_status:", err)
        } else if isAdmin > 0 { // Проверка, является ли пользователь администратором

            // Пользователь - администратор
        }else{
			 // Пользователь - не администратор
			fmt.Println("ВЫ ЗАРЕГАНЫ И МОЖЕТЕ РАБОТЬ")
		}
    }}}


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
		parts := strings.Fields(message.Text)
		if len(parts) == 3 {
			userState.FirstName = parts[0]
			userState.LastName = parts[1]
			userState.MiddleName = parts[2]
			userState.State = StateRegion

			keyboard := tgbotapi.ReplyKeyboardMarkup{
				Keyboard: [][]tgbotapi.KeyboardButton{
					{
						tgbotapi.NewKeyboardButton("Азербайджан"),
						tgbotapi.NewKeyboardButton("Армения"),
						tgbotapi.NewKeyboardButton("Грузия"),
					},
					{
						tgbotapi.NewKeyboardButton("Казахстан"),
						tgbotapi.NewKeyboardButton("Киргизия"),
						tgbotapi.NewKeyboardButton("Монголия"),
					},
					{
						tgbotapi.NewKeyboardButton("Таджикистан"),
						tgbotapi.NewKeyboardButton("Туркменистан"),
						tgbotapi.NewKeyboardButton("Узбекистан"),
					},
				},
				OneTimeKeyboard: true,
			}

			msg := tgbotapi.NewMessage(chatID, "Выберите ваш регион:")
			msg.ReplyMarkup = keyboard

			_, err := b.api.Send(msg)
			if err != nil {
				log.Println(err)
			}
		} else {
			msg := tgbotapi.NewMessage(chatID, "Пожалуйста, отправьте ФИО в правильном формате (Ф И О).")
			_, err := b.api.Send(msg)
			if err != nil {
				log.Println(err)
			}
		}
	case StateRegion:
		userState.Region = message.Text
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
		userIDStr := strconv.Itoa(message.From.ID)

		err := CreateUserAccount(userIDStr, userState.FirstName, userState.LastName, userState.MiddleName, userState.PhoneNumber, userState.Region, userState.Email)
		if err != nil {
			log.Println(err)
		}

		_, sendErr := b.api.Send(msg)
			if sendErr != nil {
				log.Println(sendErr)
			}

	}
}
